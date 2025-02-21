# find all leap years from a given year to the present


def findLeapYears(year: int) -> list:
    present = 2025

    for i in range(3):
        if year % 4 != 0:
            year += 1
            continue

        break

    # Find where to start
    leap_years = []
    for y in range(year, present, 4):
        leap_years.append(y)

    return leap_years


if __name__ == "__main__":
    print(findLeapYears(2020))
