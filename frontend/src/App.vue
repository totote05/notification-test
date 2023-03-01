<template>
  <div id="app">
    <header
      class="py-2 has-background-primary has-text-white has-text-centered"
    >
      <h1 class="is-size-3">Notification Test</h1>
    </header>
    <main class="p-5">
      <div class="columns is-desktop">
        <div class="column is-4-desktop"><message-form /></div>
        <div class="column is-8-desktop"><message-log /></div>
      </div>
    </main>
  </div>
</template>

<script>
import MessageForm from "./components/MessageForm.vue";
import MessageLog from "./components/MessageLog.vue";

export default {
  components: { MessageForm, MessageLog },
  methods: {
    onMessageSubmitted() {
      console.log("event submitted");
    },
  },
};
</script>

<style scoped>
#app {
  min-height: 100vh;
}
</style>
