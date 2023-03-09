<template>
  <q-page class="row items-center justify-evenly">
    <div class="welcome">
      <div class="welcome-content">
        <h3 class="wc-heading">Welcome to StarSearch</h3>
        <p class="wc-text">StarSearch is a simple set of API endpoints to demonstrate Golang's strength in RESTful API development. It lists StarWar movies by
          fetching data from <a href="https://swapi.dev" target="_blank">Swapi API</a>, transforms and send them as a <code>Response</code> to different endpoints.</p>
      </div>
      <div class="welcome-actions">
        <q-btn class="wa-btn" label="Demo App" icon="phone_android" @click="goTo('movie-list')" flat no-caps />
        <q-btn class="wa-btn" label="API Demo" icon="code" flat no-caps />
        <q-btn class="wa-btn" label="Documentations" icon="receipt_long" @click="goTo('docs')"  flat no-caps />
      </div>
    </div>
  </q-page>
</template>
<script lang="ts" setup>
import { useRouter } from 'vue-router';

const router = useRouter()
const goTo = (name: string) => router.push({name})
</script>
<style lang="scss" scoped>
.welcome {
  display: flex;
  flex-direction: column;
  width: 70%;
  align-items: center;

  .welcome-content {
    max-width: 692px;
    text-align: center;
  
    .wc-heading {
      margin: 12px auto;
    }
    .wc-text {
      font-size: 1.4rem;
    }
  }
  
  .welcome-actions {
    display: inline-flex;
    width: 100%;
    justify-content: center;
    column-gap: 12px;
    margin-top: 32px;
  
    .wa-btn {
      background-color: $primary;
      color: white;
      min-width: 174px;
    }
  }
}

body.screen--xs {
  .welcome-actions {
    flex-direction: column;
    row-gap: 8px;

    .wa-btn {
      // min-width: none;
      width: 100%;
    }
  }
}
</style>