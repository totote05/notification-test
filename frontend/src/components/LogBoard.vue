<template>
  <div>
    <div
      v-if="fetchErrorMessage !== ''"
      class="notification is-danger is-light has-text-centered"
    >
      We can't fetch the message logs, reason: {{ fetchErrorMessage }}
    </div>
    <div
      v-else-if="messages.length == 0"
      class="notification is-primary is-light has-text-centered"
    >
      No notification has been sent yet
    </div>
    <div v-else class="table-container">
      <table class="table is-striped is-bordered">
        <thead>
          <th>Time</th>
          <th>Category</th>
          <th>Message</th>
          <th>Channel</th>
          <th>Status</th>
          <th>User</th>
          <th>Contact</th>
          <th>Subscriptions</th>
          <th>Channels</th>
        </thead>
        <tbody>
          <tr :key="index" v-for="(message, index) in messages">
            <td>{{ message.registeredAt | formatDate }}</td>
            <td>
              <category-badge :value="message.category" />
            </td>
            <td>{{ message.message }}</td>
            <td>
              <channel-badge :value="message.response.channel" />
            </td>
            <td><status-badge :value="message.response.status" /></td>
            <td>{{ message.user.name }}</td>
            <td>
              <contact-badge :value="message" />
            </td>
            <td>
              <category-badge
                :value="category"
                :key="category"
                v-for="category in message.user.subscribed"
              />
            </td>
            <td>
              <channel-badge
                :value="channel"
                :key="channel"
                v-for="channel in message.user.channels"
              />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import CategoryBadge from "./CategoryBadge.vue";
import ChannelBadge from "./ChannelBadge.vue";
import ContactBadge from "./ContactBadge.vue";
import StatusBadge from "./StatusBadge.vue";

const dateFormatter = new Intl.DateTimeFormat("default", {
  dateStyle: "short",
  timeStyle: "long",
});

export default {
  components: { StatusBadge, CategoryBadge, ChannelBadge, ContactBadge },
  computed: {
    ...mapGetters("message", ["messages", "fetchErrorMessage"]),
  },
  filters: {
    formatDate(value) {
      return dateFormatter.format(new Date(value));
    },
  },
};
</script>
