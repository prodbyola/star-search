<template>
    <q-page class="row items-center justify-evenly">
        <h3 class="page-title">API Documentation</h3>
        <q-card class="doc-card">
            <GetMovies />
        </q-card>
    </q-page>
</template>
<script lang="ts" setup>
import GetMovies from './docs/GetMovies.vue';
</script>
