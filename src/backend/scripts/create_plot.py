import csv
import matplotlib.pyplot as plt
import numpy as np

import random

# Файл с данными
results_file = './index_test_results.csv'

# Переименование индексов для легенды
index_names = {
    'Без индексов': 'Без индексов',
    'idx_equipment_test_name': 'Простой индекс на поле name',
    'idx_equipment_test_description': 'Простой индекс на поле description',
    'idx_equipment_test_gymid': 'Составной индекс на поле gymid',
    'idx_equipment_test_name_gymid': 'Составной индекс по поля name и gymid',
    'idx_equipment_test_name_description_gymid': 'Составной индекс на поля name и description и gymid'
}

def read_csv(file_path):
    results = []
    
    # Чтение данных из CSV файла
    with open(file_path, 'r') as csvfile:
        reader = csv.reader(csvfile)
        next(reader)  # Пропустить заголовок
        for row in reader:
            num_rows = int(row[0])
            index = row[1]
            exec_time = float(row[2]) * 1000  # Преобразование времени в миллисекунды
            results.append((num_rows, index, round(exec_time, 3)))


    return results
    # indices = list(set([result[1] for result in results]))
    # colors = plt.cm.get_cmap('tab10', len(indices))

    # plt.figure(figsize=(12, 8))

    # for idx, index_name in enumerate(indices):
    #     num_rows = [result[0] for result in results if result[1] == index_name]
    #     exec_times = [result[2] for result in results if result[1] == index_name]
        
    #     plt.plot(num_rows, exec_times, 'o', label=index_names.get(index_name, index_name))  # Точки на графике
        
    #     # Нелинейная апроксимация (например, полином второй степени)
    #     z = np.polyfit(num_rows, exec_times, 2)
    #     p = np.poly1d(z)
    #     plt.plot(num_rows, p(num_rows), "--", color=colors(idx))

    # plt.xlabel('Number of Rows')
    # plt.ylabel('Execution Time (ms)')
    # plt.title('Function Execution Time with Different Indices')
    # plt.legend()
    # plt.grid(True)
    # plt.savefig('execution_times_from_csv.png')
    # plt.show()

def create_plot(results):
    indices = list(set([result[1] for result in results]))
    colors = plt.cm.get_cmap('tab10', len(indices))

    plt.figure(figsize=(12, 8))

    for idx, index_name in enumerate(indices):
        num_rows = [result[0] for result in results if result[1] == index_name]
        exec_times = [result[2] for result in results if result[1] == index_name]
        
        plt.scatter(num_rows, exec_times, color=colors(idx), label=index_name)
        
        # Квадратичная аппроксимация
        z = np.polyfit(num_rows, exec_times, 2)
        p = np.poly1d(z)
        x = np.linspace(min(num_rows), max(num_rows), 500)
        plt.plot(x, p(x), "--", color=colors(idx))

    plt.xlabel('Число записей')
    plt.ylabel('Время исполнения (с)')
    plt.title('Зависимость времени исполнения функции от числа записей для разных индексов')
    plt.legend()
    plt.grid(True)
    plt.savefig(f'execution_times{random.randint(1, 100)}.png')
    plt.show() 

# Запуск функции построения графика
# plot_results_from_csv(results_file)
def main():
    results = read_csv(results_file)
    create_plot(results)

if __name__ == "__main__":
    main()
