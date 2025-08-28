def generate_output_line(service_name, method_suffix):
    # Извлекаем название модуля из полного имени сервиса
    module_name = service_name.split('_')[1]  # берем второй элемент после разделения
    mod_variable = module_name + "Mod"  # формируем имя переменной модуля

    # Формируем строку вывода
    output_line = f'"{service_name}": {{Tok: tfbridge.MakeDataSource(mainPkg, {mod_variable}, "get{method_suffix}"), Docs: &tfbridge.DocInfo{{AllowMissing: true}}}},'
    return output_line

def process_file(input_filename):
    # Читаем строки из файла
    with open(input_filename, 'r') as file:
        lines = file.readlines()

    # Обрабатываем каждую строку
    for line in lines:
        line = line.strip()  # Удаляем пробельные символы с начала и конца строки
        if line:
            service_name = line.strip('"')  # Удаляем кавычки в начале и конце
            method_suffix = ''.join(word.capitalize() for word in service_name.split('_')[2:])  # Формируем суффикс метода
            output_line = generate_output_line(service_name, method_suffix)
            print(output_line)

# Использование функции
input_filename = 'input.txt'  # Имя файла с входными данными
process_file(input_filename)

