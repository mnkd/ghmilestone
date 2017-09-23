# ghmilestone
A command for outputting GitHub milestones and related issues.

# Usage

```
Usage:
 ghmilestone [-r repo] [--list]       : Print milestones for a repository
 ghmilestone [-r repo] [-m milestone] : Print issues for a milestone.
```

## Configuration

`$HOME/.config/ghmilestone/config.json

```json
{
  "github": {
    "access_token": "your access_token",
    "owner": "owner name",
  }
}
```

# Examples

```
$ ghmilestone -r linguist --list
* [2 - 5.0 Release](https://github.com/github/linguist/milestone/2)
```

```
$ ghmilestone -r linguist -m 2
# ISSUE
## OPEN (0)
## CLOSED (0)

# PULL REQUEST
## OPEN (0)
## CLOSED (4)
* [2006 - Use filenames as a definitive answer](https://github.com/github/linguist/pull/2006) ()
* [3278 - The grand language renaming bonanza](https://github.com/github/linguist/pull/3278) ()
* [3359 - Remove deprecated code](https://github.com/github/linguist/pull/3359) ()
* [3388 - Release v5.0.0](https://github.com/github/linguist/pull/3388) (brandonblack)
```
