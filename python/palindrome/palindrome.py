#  Check if a given string is a palindrome.
# A palindrome is a word, phrase, or sequence that reads the same backward as forward


def isPalindrome(s) -> bool:

    i, j = 0, len(s) - 1

    while i < j:
        if s[i] == s[j]:
            i += 1
            j -= 1
        else:
            return False

    return True


if __name__ == "__main__":
    print(isPalindrome("test"))
    print(isPalindrome("madam"))
