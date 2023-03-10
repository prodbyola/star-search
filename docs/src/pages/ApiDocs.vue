<template>
    <q-page class="row items-center justify-evenly">
        <h3 class="page-title">API Documentation</h3>
        <q-card class="doc-card">
            <GetMovies />
            <GetCharacters />
            <AddComment />
            <GetComments />
        </q-card>
    </q-page>
</template>
<script lang="ts" setup>
import GetMovies from './docs/GetMovies.vue';
import GetCharacters from './docs/GetCharacters.vue';
import AddComment from './docs/AddComment.vue';
import GetComments from './docs/GetComments.vue';
</script>