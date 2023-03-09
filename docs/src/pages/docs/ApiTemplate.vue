<template>
    <div class="api-detail">
                <h3 class="ad-name">{{ name }}</h3>
                <p>{{ desc }}</p>
                <p class="ad-url"><strong>URL: </strong>{{ link }}</p>
                <slot name="response">

                    
                </slot>
            </div>
</template>
<script lang="ts" setup>
defineProps<{
    name: string
    desc: string
    link: string
}>()
</script>