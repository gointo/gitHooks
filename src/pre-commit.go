// A hook script to verify what is about to be commited.
// Calles by `git commit` with no additional arguments. The hook should
// exit with non-xero status after issuing an appropriate message if it wants
// to stop the commit.
//
// To enable this hook, move the binary or call it from a shell script
// implemented in your targeted `hooks` directory.
//
package gitHooks

import (
  `log`
  `os`
  `os/exec`
)

// Check if this is the initial commit
//
func Gethead() (string, error) {
  head := os.Getenv(`HEAD`)
  cmd := exec.Command(`git`, `rev-parse`, `--verify`, head)
  err := cmd.Run()
  if err != nil {
     // Initial commit: diff against an empty tree object
     head = `4b825dc642cb6eb9a060e54bf8d69288fbee4904`
  }
  return head, err
}

// Cross platform projects tend to avoid non-ASCII filenames; prevent
// them from being added to the repository. We exploit the fact that the
// printable range starts at the space character and ends with tilde.
//
func checkNonAscii(against string) {
  nonascii := Getasciisettings()
  if nonascii != `true` {
  ascii := exec.Command(
      `git`,
      `diff`,
      `--cached`,
      `--name-only`,
      `--diff-filter=A`,
      `-z`,
      against,
      `|`,
      `LC_ALL=C`,
      `tr`,
      `-d`,
      `'[ -~]\0'`,
      `|`,
      `wc`,
      `-c`).Run()
    if ascii != nil {
      log.Fatal(`
Error: Attempt to add a non-ASCII file name.

This can cause problems if you want to work with people on other platforms.

To be portable it is advisable to rename the file.

If you know what you are doing you can disable this check using:

  git config hooks.allownonascii true`)
    }
  }
}

//
// Check for allownonascii setting.
// If you want to activate it, you need to set hooks.allownonascii using:
//
// $ git config --add hooks.allownonascii True
//
func Getasciisettings() (string) {
  res, _ := exec.Command(`git`, `config`, `--bool`, `hooks.allownonascii`).Output()
  return string(res)
}

func checkWhitspaces(against string) {
  err := exec.Command(`git`, `diff-index`, `--check`, `--cached`, against, `--`)
  if err != nil {
    log.Fatalf(`Error with Whitpaces check:\n%s`, err)
  }
}

func getUsername() string {
  return `hahaha`
}

//
// Run all pre-commit hook functions implemented
//
func main() {
  against,_ := Gethead()
  checkNonAscii(against)
  checkWhitspaces(against)
  log.Printf(`Your commit if find Sir %s! :)`, getUsername())
}
