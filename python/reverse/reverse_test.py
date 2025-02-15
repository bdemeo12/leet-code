from reverse.reverse import reverseString, reverseList

def test_reverseString():
    expected = "gnirts"
    actual = reverseString("string")

    assert expected == actual 

def test_reverseList():
    expected = [4,3,2,1]
    list = [1,2,3,4]
    actual = reverseList(list)
    print(actual)

    assert expected == actual
    