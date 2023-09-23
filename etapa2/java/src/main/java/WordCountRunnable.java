import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Semaphore;
import java.util.List;
import java.util.ArrayList;

class WordCounter implements Runnable {
	private File file;
	private int localCount = 0;
	private Semaphore mutex;

	public WordCounter(File file, Semaphore mutex) {
		this.file = file;
		this.mutex = mutex;
	}

	@Override
	public void run(){
		try (BufferedReader reader = new BufferedReader(new FileReader(file))) {
			String line;
			while ((line = reader.readLine()) != null) {
				String[] words = line.split("\\s+");
				localCount += words.length;
			}
		} catch (Exception e) {
			e.printStackTrace();
		}

		try {
			mutex.acquire();
			WordCountRunnable.count += localCount;
			mutex.release();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}

public class WordCountRunnable {
	public static int count = 0;	
	
	public static void main(String args[]){
		if (args.length != 1) {
     		       System.err.println("Usage: java WordCount <root_directory>");
   		       System.exit(1);
        	}
        	File rootDir = new File(args[0]);
        	File[] subdirs = rootDir.listFiles();
        	Semaphore mutex = new Semaphore(1);
		
		if (rootDir.exists() && rootDir.isDirectory()){
			ExecutorService exec = Executors.newFixedThreadPool(3);

			List<Thread> threads = new ArrayList<>();
			for (File file: rootDir.listFiles()) {
				if (file.isDirectory()) {
					for (File sub: file.listFiles()) {
						Runnable wordCounter = new WordCounter(sub, mutex);
						exec.execute(wordCounter);
					}
				} else {
					Runnable wordCounter = new WordCounter(file, mutex);
					exec.execute(wordCounter);
				}
			}
			exec.shutdown();
			while(!exec.isTerminated()) {}
			System.out.println(count);						
		} else {
			System.out.println(count);
		}

         }	
}
