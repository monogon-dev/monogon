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
        url = https://github.com/bazelbuild/bazel/releases/download/5.4.0/bazel-5.4.0-linux-x86_64;
        sha256 = "1w58m1brwjfwsv48fmd66inry67m4vgb3bwvwmamhdv099v183jg";
      };
      unpackPhase = ''
        true
      '';
      buildPhase = ''
        mkdir -p $out/bin
        cp $src $out/bin/.bazel-inner
        chmod +x $out/bin/.bazel-inner

        cat > $out/bin/bazel <<EOF
        #!/usr/bin/bash
        export BAZEL_REAL=$out/bin/.bazel-inner
        function get_workspace_root() {
          workspace_dir="\''${PWD}"
          while [[ "\''${workspace_dir}" != / ]]; do
            if [[ -e "\''${workspace_dir}/WORKSPACE" || -e "\''${workspace_dir}/WORKSPACE.bazel" ]]; then
              readonly workspace_dir
              return
            fi
            workspace_dir="$(dirname "\''${workspace_dir}")"
          done
          readonly workspace_dir=""
        }

        get_workspace_root
        readonly wrapper="\''${workspace_dir}/tools/bazel"
        if [ -f "\''${wrapper}" ]; then
          exec -a "\$0" "\''${wrapper}" "\$@"
        fi
        exec -a "\$0" "\''${BAZEL_REAL}" "\$@"
        EOF
        chmod +x $out/bin/bazel
      '';
      dontStrip = true;
    })
    zlib
    curl
    gcc
    binutils
    openjdk11
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
