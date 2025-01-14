{ pkgs, extraConf ? "" }: with pkgs;
let
  wrapper = pkgs.writeScript "wrapper.sh"
    ''
      # Fancy colorful PS1 to make people notice easily they're in the Monogon Nix shell.
      PS1='\[\033]0;\u/monogon:\w\007\]'
      if type -P dircolors >/dev/null ; then
        PS1+='\[\033[01;35m\]\u/monogon\[\033[01;36m\] \w \$\[\033[00m\] '
      fi
      export PS1

      # Use Nix-provided cert store.
      export NIX_SSL_CERT_FILE="${cacert}/etc/ssl/certs/ca-bundle.crt"
      export SSL_CERT_FILE="${cacert}/etc/ssl/certs/ca-bundle.crt"

      # Let some downstream machinery know we're on NixOS. This is used mostly to
      # work around Bazel/NixOS interactions.
      export MONOGON_NIXOS=yep

      # Convince rules_go to use /bin/bash and not a NixOS store bash which has
      # no idea how to resolve other things in the nix store once PATH is
      # stripped by (host_)action_env.
      export BAZEL_SH=/bin/bash

      ${extraConf}

      # Allow passing a custom command via env since nix-shell doesn't support
      # this yet: https://github.com/NixOS/nix/issues/534
      if [ ! -n "$COMMAND" ]; then
          COMMAND="bash --noprofile --norc"
      fi
      exec $COMMAND
    '';
in
(pkgs.buildFHSUserEnv {
  name = "monogon-nix";
  targetPkgs = pkgs: with pkgs; [
    git
    buildifier
    (stdenv.mkDerivation {
      name = "bazel";
      src = builtins.fetchurl {
        url = "https://github.com/bazelbuild/bazel/releases/download/7.4.1/bazel-7.4.1-linux-x86_64";
        sha256 = "0ar0qvxa3a7bms4baa84r1aq5yk53prc32k75063zrnw789h4zy9";
      };
      unpackPhase = ''
        true
      '';
      nativeBuildInputs = [ makeWrapper ];
      buildPhase = ''
        mkdir -p $out/bin
        cp $src $out/bin/.bazel-inner
        chmod +x $out/bin/.bazel-inner

        cp ${./bazel-inner.sh} $out/bin/bazel
        chmod +x $out/bin/bazel

        # Use wrapProgram to set the actual bazel path
        wrapProgram $out/bin/bazel --set BAZEL_REAL $out/bin/.bazel-inner
      '';
      dontStrip = true;
    })
    zlib
    curl
    gcc
    binutils
    openjdk21
    patch
    python3
    busybox
    niv
    google-cloud-sdk
    qemu_kvm
    swtpm
  ];
  runScript = wrapper;
})
