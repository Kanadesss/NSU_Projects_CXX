#define _GNU_SOURCE
#include <stdio.h>
#include <pthread.h>
#include <string.h>
#include <errno.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>

//  Напишите программу, которая в бесконечном цикле будет создавать поток, с
//  поточной функцией, которая выводит свой идентификатор потока и
//  завершается. Запустите. Объясните результат.

// pthread_self - const

void* my_thread() {
	printf("my_thread [%d %d %d %lu]: Hello from my_thread!\n", getpid(), getppid(), gettid(), (unsigned long)pthread_self());
	return NULL;
}

int main() {
	pthread_t tid;
	int err;
    printf("main [%d %d %d]: Hello from main!\n", getpid(), getppid(), gettid());
    while(1) {
        err = pthread_create(&tid, NULL, my_thread, NULL);
        if (err) {
            fprintf(stderr, "main: pthread_create() failed: %s\n", strerror(err));
            return EXIT_FAILURE;
        }

    //    void *ret_val;
    //    err = pthread_join(tid, &ret_val);
    //    if (err) {
    //        fprintf(stderr, "main: pthread_join() failed %s\n", strerror(err));
    //    }
    }
	return EXIT_SUCCESS;
}