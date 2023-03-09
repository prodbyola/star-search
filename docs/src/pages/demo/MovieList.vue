<template>
  <q-page class="row items-center justify-evenly">
        <h3 class="page-title">Movie List</h3>
        <div class="movie-list">
            <q-card v-for="(movie, index) in movieList" :key="index" class="movie-box">
                    <h3 class="mb-name">{{ movie.name }}</h3>
                <p class="mb-op" v-html="movie.opening_crawl" />
                <div class="mb-actions">
                    <q-btn class="mba" :label="movie.comment_count + ' Comments'" flat />
                    <q-btn class="mba" label="Characters" flat />
                </div>
            </q-card>
        </div>
    </q-page>
</template>
<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { api } from 'src/boot/axios';
import { ApiResponse } from 'src/shared';

type Movie = {
    id: number
    opening_crawl: string
    name: string
    comment_count: number
}

const movieList = ref<Movie[]>([])

onMounted(() => {
    api.get('/get-movies').then(r => {
        const resp: ApiResponse = r.data
        if(!resp.error){
            movieList.value = resp.message as Movie[]
        }
    })
})
</script>
<style lang="scss" scoped>
.movie-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    row-gap: 12px;
    width: 90%;

    .movie-box{
        width: 32%;
        padding: 16px;
        display: grid;

        .mb-name {
            margin: 20px 0;
            font-size: 2.4rem;
        }

        .mb-op {
            color: #5e5c5e;
        }

        .mb-actions {
            display: inline-flex;
            width: 100%;
            justify-content: space-between;
            align-items: center;
            color: #adadad;
            align-self: self-end;
        }
    }
}

body.screen--sm {
    .movie-box {
        width: 48%;
    }
}
body.screen--xs {
    .movie-box {
        width: 100%;
    }
}
</style>