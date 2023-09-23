import os
import sys
import time
from threading import Thread
from threading import Barrier 
from threading import Semaphore

def wc_file(filename):
    global count
    try:

        with open(filename, 'r', encoding='latin-1') as f:
            file_content = f.read()
        
        mutex.acquire()
        count += len(file_content.split())
        mutex.release()

    except FileNotFoundError:
        return 0

def wc_dir(dir_path):
    
    threads = []
    for filename in os.listdir(dir_path):
        filepath = os.path.join(dir_path, filename)

        if os.path.isfile(filepath):
            threads.append(Thread(target= wc_file, args=[filepath]))
        elif os.path.isdir(filepath):
            threads.append(Thread(target= wc_dir, args=[filepath])) 

    for t in threads:
        t.start()

    for t in threads:
        t.join() 
    

  # Chamada recursiva para diretórios

def main():
    global count

    if len(sys.argv) != 2:
        print("Usage: python", sys.argv[0], "root_directory_path")
        return
    root_path = os.path.abspath(sys.argv[1])
    wc_dir(root_path)

if __name__ == "__main__":
    count = 0
    mutex = Semaphore() 
    main()
    print(count)
