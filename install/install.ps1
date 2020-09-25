#!/usr/bin/env pwsh
# ? this installation script is based on the deno installation scripts
# Copyright 2018 the Deno authors. All rights reserved. MIT license.

$ErrorActionPreference = 'Stop'

if ($v) {
  $Version = "v${v}"
}
if ($args.Length -eq 1) {
  $Version = $args.Get(0)
}

$HostifyInstall = $env:HOSTIFY_INSTALL
$BinDir = if ($HostifyInstall) {
  "$HostifyInstall\bin"
} else {
  "$Home\.hostify\bin"
}

$HostifyZip = "$BinDir\hostify.zip"
$HostifyExe = "$BinDir\hostify.exe"

# GitHub requires TLS 1.2
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

$HostifyUri = if (!$Version) {
  $Response = Invoke-WebRequest 'https://github.com/eggheaddev/hostify-cli/releases' -UseBasicParsing
  if ($PSVersionTable.PSEdition -eq 'Core') {
    $Response.links |
      Where-Object { $_.href -like "/eggheaddev/hostify-cli/releases/download/*/xhostify-windows.zip" } |
      ForEach-Object { 'https://github.com' + $_.href } |
      Select-Object -First 1
  } else {
    $HTMLFile = New-Object -Com HTMLFile
    if ($HTMLFile.IHTMLDocument2_write) {
      $HTMLFile.IHTMLDocument2_write($Response.Content)
    } else {
      $ResponseBytes = [Text.Encoding]::Unicode.GetBytes($Response.Content)
      $HTMLFile.write($ResponseBytes)
    }
    $HTMLFile.getElementsByTagName('a') |
      Where-Object { $_.href -like "about:/eggheaddev/hostify-cli/releases/download/*/xhostify-windows.zip" } |
      ForEach-Object { $_.href -replace 'about:', 'https://github.com' } |
      Select-Object -First 1
  }
} else {
  "https://github.com/eggheaddev/hostify-cli/releases/download/${Version}/xhostify-windows.zip"
}

if (!(Test-Path $BinDir)) {
  New-Item $BinDir -ItemType Directory | Out-Null
}

Invoke-WebRequest $HostifyUri -OutFile $HostifyZip -UseBasicParsing

if (Get-Command Expand-Archive -ErrorAction SilentlyContinue) {
  Expand-Archive $HostifyZip -Destination $BinDir -Force
} else {
  if (Test-Path $HostifyExe) {
    Remove-Item $HostifyExe
  }
  Add-Type -AssemblyName System.IO.Compression.FileSystem
  [IO.Compression.ZipFile]::ExtractToDirectory($HostifyZip, $BinDir)
}

Remove-Item $HostifyZip

$User = [EnvironmentVariableTarget]::User
$Path = [Environment]::GetEnvironmentVariable('Path', $User)
if (!(";$Path;".ToLower() -like "*;$BinDir;*".ToLower())) {
  [Environment]::SetEnvironmentVariable('Path', "$Path;$BinDir", $User)
  $Env:Path += ";$BinDir"
}

Write-Output "Hostify was installed successfully to $HostifyExe"
Write-Output "Run 'hostify help' to get started"