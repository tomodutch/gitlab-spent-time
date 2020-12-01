Gitlab has no support for tracking time on commit messages. This cli application enables users to track time through commit messages according to the [proposal](https://gitlab.com/gitlab-org/gitlab/-/issues/16543#note_259395262).

# Examples

```bash
# Add time spent of 1h to issue #18
git commit -m "Spend time on #18+1h"

# Add time spent of 1h and 30m on #18 and 30m on #20
git commit -m "Spend time on #18+1h30m #20+30m"
```

**Output**

```bash
➜  test-repo git:(master) ✗ git commit -m "#18+1h #19+2h"
2020/12/01 21:05:48 Recording time for project: 22124214 issue: 18 time: 1h
2020/12/01 21:05:50 Recording time for project: 22124214 issue: 19 time: 2h
[master 8f810ad] #18+1h #19+2h
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 test3
```

# Getting started
[Download](https://github.com/TFarla/gitlab-spent-time/releases) the binary for your system.

Create a `post-commit` hook file in the `.git/hooks` directory with the following content:

```
./gitlab-hours_<your os> -project=<project> -apikey=<apikey> -repo=/path/to/git/repo
```

It's also possible to use a different gitlab url
```
./gitlab-hours_<your os> -project=<project> -apikey=<apikey> -url=https://gitlab.com/api/v4
```

Make the binary executable:

```bash
chmod +x ./gitlab-hours_<your os>
```

## Api key
Get a new personal access token [here](https://gitlab.com/-/profile/personal_access_tokens).
Make sure to grant the **api** permission so time spent requests can be made

## Project id
The project id is an integer and can be found on the repository page