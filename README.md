#Github Label Maker

Add Github labels automatically.

It's very useful when using Zenhub.io or managing a open source project that you tame the issue beast with labels. But adding multiple labels across different projects is cumbersome. BUT NOT ANYMORE.

#Deps
Please have go installed  
[GoLang](https://golang.org/doc/install)

#Install
`go get github.com/ecasilla/github_labelmaker`


###Examples
See my [example file](https://github.com/ecasilla/github_labelmaker/blob/master/labels.json)

```
[
  {"name": "bug", "color": "ffffff"},
  {"name": "feature", "color": "000000"}
]
```

##Usage
`github_labelmaker -f labels.json -u octocat -r repo`

```
OPTIONS:
   --file, -f 		A file path to your labels. i.e -f labels.json [$LABELS_FILE]
   --user, -u 		Your Github username -u octocat [$GITHUB_USER]
   --token, -t 		Your Github OAuth Token -t github_token [$GITHUB_TOKEN]
   --repo, -r 		The GITHUB repo you want to add labels to -r octocat_repo [$GITHUB_REPO]
   --help, -h		show help
   --version, -v	print the version
```

##Github Token
You will need a personal access token in order to authenticate against the Github api   
[Token Docs](https://help.github.com/articles/creating-an-access-token-for-command-line-use/)

Once you have your token your can add them to the .env file [here](https://github.com/ecasilla/github_labelmaker/blob/master/.env)

or create a env variable  
`export GITHUB_TOKEN="token"`

##GitHub Enterprise configuration
TODO:

##Export from GitHub website

Here is a snippet to be able to export Github labels from the labels page of a project

[Extract Labels Function](https://gist.github.com/ecasilla/0a70c16518aaed50a534bfa2bce750b0)

Running this code in your browsers console will output some Json ready to be save to a file for use with this tool.
