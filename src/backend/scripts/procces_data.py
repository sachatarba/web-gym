import csv
from collections import defaultdict

# Открываем исходный файл для чтения и чтения данных
with open('index_test_results.csv', mode='r') as file:
    reader = csv.reader(file)
    header = next(reader)  # Читаем заголовок

    # Создаем словарь для хранения данных по num_rows
    results_by_num_rows = defaultdict(list)

    for row in reader:
        num_rows = int(row[0])
        index = row[1]
        execution_time = float(row[2]) * 1000  # Переводим секунды в миллисекунды
        execution_time_rounded = round(execution_time, 3)  # Округляем до трех знаков после запятой

        # Добавляем данные в соответствующий список
        results_by_num_rows[num_rows].append((num_rows, index, execution_time_rounded))

# Записываем данные в отдельные файлы для каждого num_rows
for num_rows, results in results_by_num_rows.items():
    filename = f'index_test_results_{num_rows}.csv'
    with open(filename, mode='w', newline='') as file:
        writer = csv.writer(file)
        writer.writerow(header)  # Записываем заголовок
        writer.writerows(results)  # Записываем данные для текущего num_rows

    print(f"Файл {filename} успешно создан.")

print("Процесс завершен.")
