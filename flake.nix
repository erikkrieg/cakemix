{
  description = "Cakemix nix configurations";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/23.05";
    utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "utils";
    };
  };

  outputs = { utils, nixpkgs, gomod2nix, ... }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ gomod2nix.overlays.default ];
        };
        cakemix = pkgs.buildGoApplication {
          pname = "cakemix";
          version = "1.0.1";
          src = ./.;
          modules = ./gomod2nix.toml;
        };
      in
      with pkgs; {
        packages = {
          default = cakemix;
          cakemix = cakemix;
        };
        devShell = mkShell {
          buildInputs = [
            go
            gopls
            gotools
            go-tools
            gomod2nix.packages.${system}.default
            just
          ];
          shellHook = ''
            # Keep Go cache and module files in the project
            export GOPATH="$(pwd)/.go"
          '';
        };
      });
}
