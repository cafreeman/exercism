object HelloWorld {
  def hello(): String = {
    return("Hello, World!")
  }

  def hello(name: String): String = {
    return(s"Hello, $name!")
  }
}
