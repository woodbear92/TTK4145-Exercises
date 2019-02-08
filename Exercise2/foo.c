#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mtx;
// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    int j=0;
    for (j; j<999998; j++){
      pthread_mutex_lock(&mtx);
        i++;
      pthread_mutex_unlock(&mtx);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    int j=0;
    for (j; j<999999; j++){
        pthread_mutex_lock(&mtx);
        i--;
        pthread_mutex_unlock(&mtx);
    }
    return NULL;
}


int main(){
    // TODO: declare incrementingThread and decrementingThread (hint: google pthread_create)
    pthread_t incrementingThread;
    pthread_t decrementingThread;

    pthread_mutex_init (&mtx, NULL);

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);

    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);

    pthread_mutex_destroy(&mtx);

    printf("The magic number is: %d\n", i);
    return 0;
}
