<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title class="app-title" @click="goTo('home')">
          StarSearch
        </q-toolbar-title>

        <div>v1.0.0</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item @click="goTo('home')" clickable v-ripple>
          <q-item-section avatar>
            <q-icon name="home" />
          </q-item-section>
          <q-item-section>
            Home
          </q-item-section>
        </q-item>
        <q-item @click="goTo('movie-list')" clickable v-ripple>
          <q-item-section avatar>
            <q-icon name="phone_android" />
          </q-item-section>
          <q-item-section>
            Demo App
          </q-item-section>
        </q-item>
        <q-separator />
        <q-item @click="goTo('docs')" clickable v-ripple>
          <q-item-section avatar>
            <q-icon name="receipt_long" />
          </q-item-section>
          <q-item-section>
            Documentation
          </q-item-section>
        </q-item>
        <q-separator />

        
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const goTo = (name: string) => router.push({ name })
const leftDrawerOpen = ref(false)

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>
<style lang="scss" scoped>
.app-title {
  cursor: pointer;
}
</style>