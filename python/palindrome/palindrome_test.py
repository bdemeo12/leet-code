from palindrome.palindrome import isPalindrome


def test_isPalindrome_bad():
    s = "bad"
    res = isPalindrome(s)

    assert res == False


def test_isPalindrome_happy():
    s = "madam"
    res = isPalindrome(s)

    assert res == True
