# pvsadm

## Overview
IBM Power Systems Virtual Server projects deliver flexible compute capacity for Power Systems workloads.
Integrated with the IBM Cloud platform for on-demand provisioning.

This is a tool built for the IBM Power Systems Virtual Server helps managing and maintaining the resources easily.

Download the pvsadm binary from the location https://github.com/ppc64le-cloud/pvsadm/releases/
Select the latest release and download the relevant binary under the Assets section.
Run the ./pvsadm --help command to check the menu options for the pvsadm sub commands.

## pvsadm image

Sub command under the pvsadm tool to perform image related tasks like image conversion, uploading and importing into the IBM Power Systems Virtual Server instances. For more information, refer to the ./pvsadm image --help command.

The typical image workflow comprises of the following steps:

1. Download the qcow2 image.
2. Convert the downloaded qcow2 image to ova using ./pvsadm image qcow2ova command.
3. Upload the ova image to IBM Cloud Object Store Bucket using ./pvsadm image upload command.
4. Import the ova image to IBM Power Systems Virtual Server instances using ./pvsadm image import command.
