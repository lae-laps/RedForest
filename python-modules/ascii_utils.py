def println(content):
    print(content)

def println(content, color):
    color_dict = {
        'green': '121',
        'red': '196',
        'blue': '51',
        'orange': '208',
        'light-green': '118'
    }

    toprint = "\033[38;5;" + color_dict[color] + 'm' + content + "\033[m"
    print(toprint)
