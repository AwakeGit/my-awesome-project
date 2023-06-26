def read_str():
    return input('Введите строку: ').strip().split()


def roman_to_arabic(roman_number):
    roman_dict = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    decimal_num = 0
    prev_value = 0

    for char in reversed(roman_number):
        current_value = roman_dict[char]

        if current_value >= prev_value:
            decimal_num += current_value
        else:
            decimal_num -= current_value

        prev_value = current_value

    return decimal_num


def decimal_to_roman(decimal_num):
    roman_dict = {1000: 'M', 900: 'CM', 500: 'D', 400: 'CD', 100: 'C', 90: 'XC',
                  50: 'L', 40: 'XL', 10: 'X', 9: 'IX', 5: 'V', 4: 'IV', 1: 'I'}
    roman_num = ''

    for value, symbol in roman_dict.items():
        while decimal_num >= value:
            roman_num += symbol
            decimal_num -= value

    return roman_num


def calculate_rim(num_one, func, num_two):
    num_one = roman_to_arabic(num_one)
    num_two = roman_to_arabic(num_two)
    result = calculate_arab(num_one, func, num_two)
    return decimal_to_roman(result)


def calculate_arab(num_one, func, num_two):
    num_one = int(num_one)
    num_two = int(num_two)

    match func:
        case "+":
            return num_one + num_two
        case "-":
            return num_one - num_two
        case "*":
            return num_one * num_two
        case "/":
            return num_one // num_two


def check(arg):
    rim = {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
    arab = {"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
    operation = {"+", "-", "/", "*"}

    if len(arg) == 3:
        num_one, func, num_two = arg
        if num_one in rim and num_two in rim and func in operation:
            return "Проверка пройдена, Римские числа\n" + calculate_rim(num_one, func, num_two)

        elif num_one in arab and num_two in arab and func in operation:
            return "Проверка пройдена, Арабские числа\n" + str(calculate_arab(num_one, func, num_two))

        else:
            return "Ошибка: Используются одновременно разные системы счисления."

    elif len(arg) > 3:
        return "Ошибка: Формат математической операции не соответствует требованиям " \
               "- должно быть два операнда и один оператор (+, -, /, *)."

    else:
        return "Ошибка: Строка не является математической операцией."


def main():
    txt = read_str()
    print(check(txt))


if __name__ == '__main__':
    main()
