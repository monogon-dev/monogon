## Smalltown IAM

There are 4 kinds of elements in Smalltown's Authorization system
* Identities
    * User
    * Key
    * Module
* Objects
    * Key
    * Secret
    * Module
* Policies
* Permissions

### Identity
Identities represent an actor that can execute **actions** like editing or interacting with an object.

Identities possess **permissions** and **properties** which can be accessed by policies.

### Objects
Objects are things that can be interacted with like keys, secrets or modules.

Each object has a **policy** that handles authorization of **actions** performed on it.

When an object is created a default policy is attached which forwards all decisions to the global policy.
For the first iteration of the system this policy will not be modifiable.

**WARNING**: by modifying a policy, an object could become inaccessible!

### Permissions

Permissions can be assigned to an identity.

| Property | Description | Example |
|----------|-------------|---------|
| Allowed Action | Regex specifying the allowed actions | key:meta:edit |
| Object | Regex specifying the objects this affects | keys:* |
| Multisig | Number of approvals required | 2 |

Optionally a permission can have a multisig flag that requires N approvals from identities with the same permission.

### Policies

Policies guard actions that are performed on an object.

By default a global policy governs all objects and global actions using an AWS IAM like model. 

Potentially a dynamic model using attachable policies could be implemented in the future to allow
for highly custom models.

A potential graphical representation of a future policy:

![graphical representation](https://i.imgur.com/CuURwjr.png)

### Global Default Ruleset

This default global policy defines an AWS IAM like permission system.

The following actions are implemented on objects:

| Category | Action | Description | Note |
|----------|-------------|---------|---------|
| Object | object:view | Allow to view the object | Cannot be scripted using the policy builder |
| Object | object:delete | Allow to delete the object |
| Object | object:attach:normal | Allow to attach the object to a module slot |
| Object | object:attach:exclusive | Allow to attach the object to an exclusive module slot |
| Object | object:policy:view | Allow to view the object's attached policy |
| Object | object:policy:edit | Allow to edit the object's attached policy |
| Object | object:audit:view | Allow to view the object's audit log |
| Object:Key | key:sign:eddsa | Allow to sign using the key |
| Object:Key | key:sign:ecdsa | Allow to sign using the key |
| Object:Key | key:sign:rsa | Allow to sign using the key |
| Object:Key | key:encrypt:rsa | Allow to encrypt using the key |
| Object:Key | key:encrypt:des | Allow to encrypt using the key |
| Object:Key | key:encrypt:3des| Allow to encrypt using the key |
| Object:Key | key:encrypt:aes | Allow to encrypt using the key |
| Object:Key | key:decrypt:rsa | Allow to decrypt using the key |
| Object:Key | key:decrypt:des | Allow to decrypt using the key |
| Object:Key | key:decrypt:3des| Allow to decrypt using the key |
| Object:Key | key:decrypt:aes | Allow to decrypt using the key |
| Object:Key | key:auth:hmac | Allow to auth messages using the key |
| Object:Secret | secret:reveal | Allow to reveal a secret to the identity |
| Object:Module | module:update | Allow to update a module's bytecode | Updates verify the module signature
| Object:Module | module:config | Allow to configure a module | Assigning objects to slots requires additional permissions on that object
| Object:Module | module:call:* | Allow to call a function of the module | Function names are defined in the module and vary between modules

The following actions are implemented globally:

| Category | Action | Description | Note |
|----------|-------------|---------|---------|
| Object | g:key:generate | Allow to generate a key |
| Object | g:key:import | Allow to import a key |
| Object | g:secret:import | Allow to import a secret |
| Object | g:module:install | Allow to install a module |
| Object | g:user:create | Allow to create a user |
| Object | g:user:permission_remove | Allow to create a user | **Privilege Escalation Risk**: Recommend Multisig
| Object | g:user:permission_add | Allow to create a user | **Privilege Escalation Risk**: Recommend Multisig
| Object | g:cluster:view | Allow to view cluster nodes
| Object | g:cluster:add | Allow to add a node to the cluster | **Dangerous**: Recommend Multisig
| Object | g:cluster:remove | Allow to remove a node from the cluster | **Dangerous**: Recommend Multisig
| Object | g:config:edit | Allow to edit the global config

