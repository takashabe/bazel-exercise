def sample(name, word="hello", **kwargs):
  native.genrule(
    name = name,
    word = word,
    cmd = "say {}".format(word),
    **kwargs
  )
