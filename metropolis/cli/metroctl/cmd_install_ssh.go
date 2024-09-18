package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	xssh "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/go/net/ssh"
	"source.monogon.dev/osbase/fat32"
)

var sshCmd = &cobra.Command{
	Use:     "ssh --disk=<disk> <target>",
	Short:   "Installs Metropolis on a Linux system accessible via SSH.",
	Example: "metroctl install --bundle=metropolis-v0.1.zip --takeover=takeover ssh --disk=nvme0n1 root@ssh-enabled-server.example",
	Args:    cobra.ExactArgs(1), // One positional argument: the target
	RunE: func(cmd *cobra.Command, args []string) error {
		user, address, err := parseSSHAddr(args[0])
		if err != nil {
			return err
		}

		diskName, err := cmd.Flags().GetString("disk")
		if err != nil {
			return err
		}

		if len(diskName) == 0 {
			return fmt.Errorf("flag disk is required")
		}

		var authMethods []xssh.AuthMethod
		if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
			defer aconn.Close()
			a := agent.NewClient(aconn)
			authMethods = append(authMethods, xssh.PublicKeysCallback(a.Signers))
		} else {
			log.Printf("error while establishing ssh agent connection: %v", err)
			log.Println("ssh agent authentication will not be available.")
		}

		if term.IsTerminal(int(os.Stdin.Fd())) {
			authMethods = append(authMethods,
				xssh.PasswordCallback(func() (string, error) {
					fmt.Printf("%s@%s's password: ", user, address)
					b, err := terminal.ReadPassword(syscall.Stdin)
					if err != nil {
						return "", err
					}
					fmt.Println()
					return string(b), nil
				}),
				xssh.KeyboardInteractive(func(name, instruction string, questions []string, echos []bool) ([]string, error) {
					answers := make([]string, 0, len(questions))
					for i, q := range questions {
						fmt.Print(q)
						if echos[i] {
							if _, err := fmt.Scan(&questions[i]); err != nil {
								return nil, err
							}
						} else {
							b, err := terminal.ReadPassword(syscall.Stdin)
							if err != nil {
								return nil, err
							}
							fmt.Println()
							answers = append(answers, string(b))
						}
					}
					return answers, nil
				}),
			)
		} else {
			log.Println("stdin is not interactive. password authentication will not be available.")
		}

		cl := ssh.DirectClient{
			Username:    user,
			AuthMethods: authMethods,
		}

		ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
		conn, err := cl.Dial(ctx, address, 5*time.Second)
		if err != nil {
			return fmt.Errorf("error while establishing ssh connection: %w", err)
		}

		params, err := makeNodeParams()
		if err != nil {
			return err
		}
		rawParams, err := proto.Marshal(params)
		if err != nil {
			return fmt.Errorf("error while marshaling node params: %w", err)
		}

		const takeoverTargetPath = "/root/takeover"
		const bundleTargetPath = "/root/bundle.zip"
		bundle, err := external("bundle", "_main/metropolis/node/bundle.zip", bundlePath)
		if err != nil {
			return err
		}
		takeover, err := external("takeover", "_main/metropolis/cli/takeover/takeover_bin_/takeover_bin", bundlePath)
		if err != nil {
			return err
		}

		barUploader := func(r fat32.SizedReader, targetPath string) {
			bar := progressbar.DefaultBytes(
				r.Size(),
				targetPath,
			)
			defer bar.Close()

			proxyReader := progressbar.NewReader(r, bar)
			defer proxyReader.Close()

			if err := conn.Upload(ctx, targetPath, &proxyReader); err != nil {
				log.Fatalf("error while uploading %q: %v", targetPath, err)
			}
		}

		log.Println("Uploading required binaries to target host.")
		barUploader(takeover, takeoverTargetPath)
		barUploader(bundle, bundleTargetPath)

		// Start the agent and wait for the agent's output to arrive.
		log.Printf("Starting the takeover executable at path %q.", takeoverTargetPath)
		_, stderr, err := conn.Execute(ctx, fmt.Sprintf("%s -disk %s", takeoverTargetPath, diskName), rawParams)
		stderrStr := strings.TrimSpace(string(stderr))
		if stderrStr != "" {
			log.Printf("Agent stderr: %q", stderrStr)
		}
		if err != nil {
			return fmt.Errorf("while starting the takeover executable: %w", err)
		}

		return nil
	},
}

func parseAddrOptionalPort(addr string) (string, string, error) {
	if addr == "" {
		return "", "", fmt.Errorf("address is empty")
	}

	idx := strings.LastIndex(addr, ":")
	// IPv4, DNS without Port.
	if idx == -1 {
		return addr, "", nil
	}

	// IPv4, DNS with Port.
	if strings.Count(addr, ":") == 1 {
		return addr[:idx], addr[idx+1:], nil
	}

	// IPv6 with Port.
	if addrPort, err := netip.ParseAddrPort(addr); err == nil {
		return addrPort.Addr().String(), fmt.Sprintf("%d", addrPort.Port()), nil
	}

	// IPv6 without Port.
	if addr, err := netip.ParseAddr(addr); err == nil {
		return addr.String(), "", nil
	}

	return "", "", fmt.Errorf("failed to parse address: %q", addr)
}

func parseSSHAddr(s string) (string, string, error) {
	user, rawAddr, ok := strings.Cut(s, "@")
	if !ok {
		return "", "", fmt.Errorf("SSH user is mandatory")
	}

	addr, port, err := parseAddrOptionalPort(rawAddr)
	if err != nil {
		return "", "", err
	}
	if port == "" {
		port = "22"
	}

	return user, net.JoinHostPort(addr, port), nil
}

func init() {
	sshCmd.Flags().String("disk", "", "Which disk Metropolis should be installed to")
	sshCmd.Flags().String("takeover", "", "Path to the Metropolis takeover binary")

	installCmd.AddCommand(sshCmd)
}
