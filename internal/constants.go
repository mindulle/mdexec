package internal

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/mdexec"

// github username and repo in order to fetch snippets.
var Owner = "hoehwa"
var Repository = "executable"

// urls for hosted websites.
const Baseurl = "https://github.com/hoehwa/snippet/tree/main"
