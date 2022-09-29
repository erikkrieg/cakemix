{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = [ pkgs.buildPackages.go_1_18 pkgs.buildPackages.gopls pkgs.buildPackages.just ];
}

