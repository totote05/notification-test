<template>
  <div>
    <h3 class="is-size-4">Notification logs</h3>
    <hr />
    <div class="has-text-centered" v-if="loading">
      Loading...
      <progress class="progress is-info"></progress>
    </div>
    <log-board v-else />
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import LogBoard from "./LogBoard.vue";

export default {
  components: { LogBoard },
  computed: {
    ...mapGetters("message", ["loading"]),
  },
  methods: {
    ...mapActions("message", ["fetchMessages"]),
  },
  mounted() {
    this.fetchMessages();
  },
};
</script>
