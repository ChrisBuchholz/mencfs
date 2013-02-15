MEncFS - Manage EncFS
=====================

MEncFS is a manager for the filesystem encryptor
[EncFS](http://www.arg0.net/encfs). MEncFS aims to make it trivial to mount,
unmount and automount EncFS encrypted folders on Mac OS X and Linux.

At the moment, MEncFS only runs on Mac OS X, but Linux support is coming.
It has been tested on Mac OS X 10.8.

Feel free to come with input and pull request.

## Installation

MEncFS depends on EncFS, which can easily be installed via
[Homebrew](http://mxcl.github.com/homebrew/) on Mac OS X, and is available in
most Linux distro's package repository.

To install MEncFS, a proper [Go](http://golang.org/) setup is required. You
can read more about installing and setting up Go on your system
[here](http://golang.org/doc/install).

When you are ready to install MEncFS, execute the following command in your
terminal of choise.

    go get github.com/ChrisBuchholz/mencfs

That's it. Try to execute `mencfs` in a terminal. If you have set up Go
correctly, you should see information in your terminal about how to use MEncFS.
If not, refer to the
[Go installation instructions](http://golang.org/doc/install) and when you
you are set, execute the command again.

## Getting Started

### Configuration

When MEncfs has been installed on your system, it's time to configure it.
Run this in a terminal:

    mencfs generate

This will generate a new MEncFS configuration file *~/.mencfs*. This is starting
point that you have to alter to use MEncFS. Open the file in your favovorite
editor.

The configuration file is defined in a format where each line descripes an
EncFS encrypted folder, the name it should have when mounted and the title
which the encryption password stored in your systems keychain is labeled with.
That's right. MEncFS don't want to handle your passwords. It relies on your
system-wide keychain. [Here](https://support.apple.com/kb/PH7282)
is a guide describing how to add a password to your Mac OS X keychain. If
you are running on another platform, you can probably Google it.

Before continuing, make sure you already have an EncFS encrypted folder. If not,
you can encrypt a folder with the following command:

    mencfs encrypt ~/my_folder

You will get asked a few questions. Just follow the prompt till it's
done. 

Add a new password to your keychain using the same password that you used to
encrypt *~/my_folder* in the previous step. Remember the label your give it.

Now alter *~/.mencfs* so it looks like this:

    ~/my_folder		my_folder		%password_label%

Ensure that *%password_label%* is the same label that you gave the password
you just added to your keychain.

### Usage

Now that your configuration file is set up, you are ready to manage your 
EncFS encrypted folder.

In a terminal, execute

    mencfs mount

This will, on Mac OS X, mount *~/my_folder* to */Volumes/my_folder*. It
should pop up in the sidebar in Finder, ready for you to put secret stuff in.

To unmount the decrypted folder again, type the following in a terminal
and press return:

    mencfs umount

That's it!

You can add as many encrypted folders to your configuration file as you'd like,
and MEncFS will be happy to manage them all for you!

## To Do

* Automount
* Linux support
