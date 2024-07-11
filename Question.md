# Question: Word Break II (Hard)
## Description:

Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences. Note that the same word in the dictionary may be reused multiple times in the segmentation.

```python
def word_break(s: str, wordDict: list) -> list:
    pass
#  Example
assert word_break("catsanddog", ["cat", "cats", "and", "sand", "dog"]) == [
    "cat sand dog",
    "cats and dog"
]
assert word_break("pineapplepenapple", ["apple", "pen", "applepen", "pine", "pineapple"]) == [
    "pine apple pen apple",
    "pineapple pen apple",
    "pine applepen apple"
]

assert word_break("catsandog", ["cats", "dog", "sand", "and", "cat"]) == []
```