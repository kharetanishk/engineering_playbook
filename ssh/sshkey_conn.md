# Generating the Ssh keys (ssh access)

### Step 1 — Generate a new SSH key pair on your SYSTEM

On your system  **terminal** (not on the VM), run:

```
ssh-keygen-t ed25519-C"macbook-air-m5"
```

-C is for comment 

we can use -f also to have specific name of the pub and private keys file 

### `-f`

Specifies the **file name/path** where the key pair will be saved.

```
-f ~/.ssh/name_vm
```

creates:

```
~/.ssh/name_vm       ← Private key
~/.ssh/name_vm.pub   ← Public key
```

It will ask:

```
Enter file in which to save the key
```

Just press **Enter** to use the default location.

Then it may ask for a passphrase:

- Press Enter twice for no passphrase (easier).
- Or set one if you want extra security.

NOW IN YOUR SYSTEM YOU WILL BE HAVING 2 SSH Keys 

one private 
one public 

you can see the keys in the folder .ssh 

from your root 
run 

ls -a

you will see .ssh folder 

then do → cd .ssh where you will be seeing all the public and private keys 

### Step 2 — Copy the public key

On your Mac, run:

```
cat ~/.ssh/filename.pub
```

You'll get a single long line like:

```
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... macbook-air-m5
```

Copy the **entire line**.

### 

### Step 3 — Add it to the VM

Go back to your VM terminal and run:

```
nano ~/.ssh/authorized_keys
```

Scroll to the bottom, paste the entire public key on a **new line**, then:

- `Ctrl + O` → Enter (save)
- `Ctrl + X` (exit)

TO know the ipv4 of our vm 
run this command 
curl -4 ifconfig.me

and one thing initially you will need vm root access command and password also 

now you can ssh  like 
go to your root of the system (your laptop ) 

command options - 

ssh -i ~/.ssh/name_vm root@72.60.222.26. | here i. have named the filename , not using default filenaming like id_ed…

### Option 2 (Recommended)

Create `~/.ssh/config`:

```
Host any_name
    HostName ip
    User root
    IdentityFile ~/.ssh/name_vm
```

Then simply:

```
ssh any_name
```

### 

summary 

generate ssh keys , give comments to it and custom filename  | → in your laptop 

then ssh into you vm using password and in the ssh folder create an authorized keys file and paste the public keys there 

then in your laptop ssh config  → paste code like host  hostname ip user root identitiyfile 

so that you can ssh like ssh hostname
