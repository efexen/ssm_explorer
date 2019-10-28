# AWS SSM Explorer

Explore your AWS SSM parameters visually by treating forward slashes as an organisational structure.

For example given the following SSM keys:

```
/my_project/db/user
/my_project/db/password
/my_project/api/key
/other_project/db/user
```

This will render the following tree:

```
[root]
  my_project
    db
      user
      password
    api
      key
  other_project
    db
      user
```

Only keys are rendered and all values are excluded for security purposes.

### Usage

Build the executable with

```
go build
```

Then simply

```
aws ssm describe-parameters | ./ssm_explorer
```

**note** non OSX users you will have to edit the command to launch browser.
