package main

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/spf13/cobra"

	"source.monogon.dev/build/toolbase"
)

var (
	flagDeep               bool
	flagMirrorBucketName   string
	flagMirrorBucketSubdir string
)

var rootCmd = &cobra.Command{
	Use:          "mirror",
	Short:        "Developer/CI tool to make sure our RPM mirror for the sandboxroot is up to date",
	SilenceUsage: true,
}

// ourMirrorURL returns a fully formed URL-string to our mirror (as defined by
// flags), optionally appending the given parts as file path parts.
func ourMirrorURL(parts ...string) string {
	var u url.URL
	u.Scheme = "https"
	u.Host = "storage.googleapis.com"

	path := []string{
		flagMirrorBucketName,
		flagMirrorBucketSubdir,
	}
	path = append(path, parts...)
	u.Path = filepath.Join(path...)
	return u.String()
}

// progress is used to notify the user about operational progress.
func progress(done, total int) {
	fmt.Printf("%d/%d files done...\r", done, total)
}

func checkMirrorURLs(rpms []*rpmDef) error {
	log.Printf("Checking all RPMs are using our mirror...")
	allCorrect := true
	for _, rpm := range rpms {
		urls := rpm.urls

		haveOur := false
		haveExternal := false
		for _, u := range urls {
			if strings.HasPrefix(u.String(), ourMirrorURL()) {
				haveOur = true
			} else {
				haveExternal = true
			}
			if haveOur && haveExternal {
				break
			}
		}
		if !haveOur {
			allCorrect = false
			log.Printf("RPM %s does not contain our mirror in its URLs", rpm.name)
		}
		if !haveExternal {
			allCorrect = false
			log.Printf("RPM %s does not contain any upstream mirror in its URLs", rpm.name)
		}
	}
	if !allCorrect {
		return fmt.Errorf("some RPMs have incorrect mirror urls")
	}
	return nil
}

func getRepositoriesBzl() string {
	ws, err := toolbase.WorkspaceDirectory()
	if err != nil {
		log.Fatalf("Failed to figure out workspace location: %v", err)
	}
	return filepath.Join(ws, "third_party/sandboxroot/repositories.bzl")
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check that everything is okay (without performing actual mirroring)",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := getRepositoriesBzl()
		rpms, err := getBazelDNFFiles(path)
		if err != nil {
			return fmt.Errorf("could not get RPMs from %s: %w", path, err)
		}

		if err := checkMirrorURLs(rpms); err != nil {
			return err
		}

		if !flagDeep {
			log.Printf("Checking if all files are present on mirror... (use --deep to download and check hashes)")
		} else {
			log.Printf("Verifying contents of all mirrored files...")
		}

		hasAll := true
		for i, rpm := range rpms {
			has, err := rpm.validateOurs(cmd.Context(), flagDeep)
			if err != nil {
				return fmt.Errorf("checking %s failed: %w", rpm.name, err)
			}
			if !has {
				log.Printf("Missing %s in mirror", rpm.name)
				hasAll = false
			}
			progress(i+1, len(rpms))
		}
		if !hasAll {
			return fmt.Errorf("some packages missing in mirror, run `mirror sync`")
		} else {
			log.Printf("All good.")
		}

		return nil
	},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Mirror all missing dependencies",
	Long: `
Check existence of (or download and verify when --deep) of every file in our
mirror and upload it if it's missing. If an upload occured, a full re-download
will be performed for verification.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		path := getRepositoriesBzl()
		rpms, err := getBazelDNFFiles(path)
		if err != nil {
			return fmt.Errorf("could not get RPMs from %s: %w", path, err)
		}

		if err := checkMirrorURLs(rpms); err != nil {
			return err
		}

		client, err := storage.NewClient(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "could not find default credentials") {
				log.Printf("Try running gcloud auth application-default login --no-browser")
			}
			return fmt.Errorf("could not build google cloud storage client: %w", err)
		}
		bucket := client.Bucket(flagMirrorBucketName)

		if !flagDeep {
			log.Printf("Checking for any missing files...")
		} else {
			log.Printf("Verifying all files and uploading if missing or corrupted...")
		}

		for i, rpm := range rpms {
			has, err := rpm.validateOurs(ctx, flagDeep)
			if err != nil {
				return err
			}
			if !has {
				log.Printf("Mirroring %s...", rpm.name)
				if err := rpm.mirrorToOurs(ctx, bucket); err != nil {
					return err
				}
				log.Printf("Verifying %s...", rpm.name)
				has, err = rpm.validateOurs(ctx, true)
				if err != nil {
					return err
				}
				if !has {
					return fmt.Errorf("post-mirror validation of %s failed", rpm.name)
				}
			}
			progress(i+1, len(rpms))
		}

		log.Printf("All good.")
		return nil
	},
}

func main() {
	rootCmd.PersistentFlags().StringVar(&flagMirrorBucketName, "bucket_name", "monogon-infra-public", "Name of GCS bucket to mirror into.")
	rootCmd.PersistentFlags().StringVar(&flagMirrorBucketSubdir, "bucket_subdir", "mirror", "Subpath in bucket to upload data to.")
	rootCmd.PersistentFlags().BoolVar(&flagDeep, "deep", false, "Always download files fully during check/sync to make sure the SHA256 matches.")
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.Execute()

}
