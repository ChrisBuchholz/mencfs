MEncFS - Manage EncFS
=====================

MEncFS is a manager for the filesystem encryptor
[EncFS](http://www.arg0.net/encfs).

MEncFS aims to make it trivial to mount, unmount and automount EncFS encrypted
folders on Mac OS X and Linux.

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

    go get https://github.com/ChrisBuchholz/mencfs

That's it. Try to execute `mencfs` in a terminal. If you have set up Go
correctly, you should see information in your terminal about how to use MEncFS.
If not, refer to the
[Go installation instructions](http://golang.org/doc/install).

## Usage

### Configuration

When MEncfs has been installed on your system, it's time to configure it.
Execute this in a terminal:

    mencfs generate

This will generate a new MEncFS configuration file ~/.mencfs. This is starting
point, but you have to alter it to use MEncFS. Open the file in your favovorite
editor.

The configuration file is defined in a format where each line descripes an
EncFS encrypted folder, the name it should have when mounted and the title
which the encryption password stored in your systems keychain is labeled with.
That's right. MEncFS don't want to handle your passwords. It relies on your
system-wide keychain technology. [Here](https://support.apple.com/kb/PH7282)
is a guide descriping how to add a password to your Mac OS X keychain. If
you are running on another platform, you can probably Google it.

Before continuing, make sure you already have a EncFS encrypted folder. If not,
you can create an encrypted folder ~/encrypted_folder like this:

    encfs ~/encrypted_folder ~/decrypted_folder

EncFS will ask you for different settings. Just follow the prompt till it's
done.

Make sure that ~/encrypted_folder is empty before you do this or else the stuff
already inside the folder won't get encrypted.

Now the two folder are synced, and if you put stuff inside ~/decrypted_folder,
it will get encrypted and put inside ~/encrypted_folder

Now unmount ~/decrypted_folder again. On Mac OS X, this can be done by
executing this command in your terminal:

    umount ~/decrypted_folder

Add a new password to your keychain with the same password you used to encrypt
~/encrypted_folder with. Remember the label your give it.

Now alter ~/.mencfs so it looks like this:

    ~/encrypted_folder		decrypted		<password_label>

Make sure that *<password_label>* is the same label that you gave the password
your created in your keychain.

### Usage

Now your configuration file is set up, and you are ready to manager your 
EncFS encrypted folder.

In a terminal, execute

    mencfs mount

This will on Mac OS X mount ~/encrypted_folder to /Volumes/decrypted. It should
pop up in the sidebar of Finder, ready for you to put secret stuff in.

To unmount your encrypted folders again, type the following in a terminal
and press return:

    mencfs umount

That's it!

You can add as many encrypted folders to your configuration file, and MEncFS
will be happy to manage them for you!

## To Do

* Automount
* Linux support
