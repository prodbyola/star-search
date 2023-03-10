<template>
    <div class="api-detail">
        <h3 class="ad-name">{{ name }}</h3>
        <p>{{ desc }}</p>
        <p class="ad-url">URL: <span class=ad-link>{{ link }}</span></p>
        <p class="adr-head">Method: <span class="adr-span">{{ method }}</span></p>
        <template v-if="hasParams">
            <p class="adr-head">Parameters:</p>
            <div class="ad-resp-indent">
                <slot name="parameters"></slot>
            </div>
        </template>
        <template v-if="hasQueries">
            <p class="adr-head">Queries:</p>
            <div class="ad-resp-indent">
                <slot name="queries"></slot>
            </div>
        </template>
        <template v-if="hasParams">
            <p class="adr-head">Request Body:</p>
            <div class="ad-resp-indent">
                <slot name="body"></slot>
            </div>
        </template>
        <p class="adr-head">Response:</p>
        <div class="ad-resp-indent">
            <div class="sig">
                <p class="resp-sig">error: <code class="light">boolean</code></p>
                <p>
                    Indicates whether the request was successful or not. If <code>true</code>, it means the request is not successful because of an error. The error message
                    would be available in the <code>message</code> property of the response. If <code>false</code>, it means the request is successful and the resource would be 
                    <code>message</code> property of the response.
                </p>
            </div>
            <div class="sig">
                <p class="resp-sig">code: <code class="light">number</code></p>
                <p>
                    The response code. <code>200</code> usually means the request is successful. If there's an error, the error code would be found here.
                </p>
            </div>
            <slot name="response"></slot>
        </div>
    </div>
</template>
<script lang="ts" setup>
defineProps<{
    name: string
    desc: string
    link: string
    method: string
    hasParams?: boolean
    hasQueries?: boolean
}>()
</script>