package utils


func WelcomeUser(fullName string) string {
  if fullName != "" {
      return "Welcome Mr/Mrs " + fullName
  }
  return ""
}
