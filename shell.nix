# If you're on NixOS, use me! `nix-shell --pure`.
{ sources ? import third_party/nix/sources.nix }:
let
    pkgs = import sources.nixpkgs {};
in
(import third_party/nix/env.nix { inherit pkgs; }).env
