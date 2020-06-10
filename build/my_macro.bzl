def sample(name, word="hello", **kwargs):
  native.genrule(
    name = name,
    word = word,
    outs = ["small_" + word],
    cmd = "say {}".format(word),
    **kwargs
  )
