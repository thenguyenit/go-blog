####T-Rex git model

After 3 years using Git for version control, I've found best practices model for myself. It's good for branching, deploy management. 

<center>
![Git Model](http://t-rex.click/storage/article/2017/git-process/git-model.jpeg)
</center>

<br>

####Branch information:

**develop:** going to run on develop environment http://dev.t-rex.click/

**release-*:** going to run on test environment http://release.t-rex.click/

**master:** going to run on staging environment http://staging.t-rex.click/ and production environment [http://t-rex.click/](http://t-rex.click/)

<br>

####Feature branch

May branch checkout from: **develop**

```
$ git checkout -b JIRA-ID origin/develop
```

After finished must merge back to **develop** via merge request

Merge strategy is: Merge commit _--no-ff_

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/pull-request--no-ff.png)

Merge strategy:

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/merge-without-ff.jpeg)

<br>

####Release branch

May branch checkout from: **develop**

After finished testing must merge back to **master** via merge request

Branch naming convention: **release-***

<br>

####Finishing release branch

When the state of the release branch is ready to become a real release

Must merge back to master and create a tag version:
```
$ git checkout master
Switched to branch 'master'
$ git merge --no-ff release-1.2
Merge made by recursive.
(Summary of changes)
$ git tag -a 1.2
```

<br>

####Hotfix branch

Branch name convention: **hotfix-***

May branch checkout from: **master**

```
$ git checkout -b hotfix-* origin/master
```

After finished must merge back to **master** and **develop**

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/hotfix-branches.jpeg)


Create a tag version:

```
$ git tag -a 0.6 -m "Version 0.6"
$ git push origin --tags
```

QA team going to deploy to staging environment for testing the deploy to real environment

<br>

####Undo

**git reset** Commit level
```
git reset [--hard|–soft] HEAD~2
```
Common use cases:
- **Discard commits** in a **private branch** or throw away uncommited changes

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/reset-commit.png)

**git reset** File level
```
git reset HEAD~2 foo.php
```
Common use cases:
- This command will fetch the version of foo.php in the 2nd-to-last commit and stage it for the next commit
 
![Merge pull request](http://t-rex.click/storage/article/2017/git-process/reset-file.png)

**git checkout** Commit level
``` 
git checkout HEAD~2
``` 
Common use cases:
- Switch between branches or inspect old snapshots <br>

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/checkout-commit.png)

```
git checkout HEAD~2 foo.php
```
Common use cases:
- Discard changes in the working directory
- Checking out a file is similar to using git reset with a file path, except it updates the working directory instead of the stage

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/checkout-file.png)    

**git revert** Commit level
```
git revert HEAD~2
```
Common use cases:
- Undo commits in a **public branch**

![Merge pull request](http://t-rex.click/storage/article/2017/git-process/revert-commit.png)

<br>

**References:**

- [A successful Git branching model](http://nvie.com/posts/a-successful-git-branching-model/)
- [Git tutorial: using branches](https://www.atlassian.com/git/tutorials/using-branches)