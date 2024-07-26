<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div v-if="!isLoaded">
    <v-progress-linear
      indeterminate
      color="primary darken-2"
    ></v-progress-linear>
  </div>
  <div v-else>
    <EditDialog
      v-model="editDialog"
      :save-button-text="$t('save')"
      :title="$t('editRunners')"
      :max-width="500"
      @save="loadItems"
    >
      <template v-slot:form="{ onSave, onError, needSave, needReset }">
        <RunnerForm
          :project-id="projectId"
          :item-id="itemId"
          @save="onSave"
          @error="onError"
          :need-save="needSave"
          :need-reset="needReset"
        />
      </template>
    </EditDialog>

    <ObjectRefsDialog
      object-title="runner"
      :object-refs="itemRefs"
      :project-id="projectId"
      v-model="itemRefsDialog"
    />

    <YesNoDialog
      :title="$t('deleteRunner')"
      :text="$t('askDeleteRun')"
      v-model="deleteItemDialog"
      @yes="deleteItem(itemId)"
    />

    <v-toolbar flat >
      <v-app-bar-nav-icon @click="showDrawer()"></v-app-bar-nav-icon>
      <v-toolbar-title>{{ $t('runners2') }}</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>

    <v-data-table
      :headers="headers"
      :items="items"
      hide-default-footer
      class="mt-4"
      :items-per-page="Number.MAX_VALUE"
    >
      <template v-slot:item.project="{ item }">
        {{ (projects.find((x) => x.id === item.project_id) || {name: '-'}).name }}
      </template>
      <template v-slot:item.inventory="{ item }">
        {{ (inventories.find((x) => x.id === item.inventory_id) || {name: '-'}).name }}
      </template>
      <template v-slot:item.actions="{ item }">
        <div style="white-space: nowrap">
          <v-btn
            icon
            class="mr-1"
            @click="askDeleteItem(item.id)"
          >
            <v-icon>mdi-delete</v-icon>
          </v-btn>

          <v-btn
            icon
            class="mr-1"
            @click="editItem(item.id)"
          >
            <v-icon>mdi-pencil</v-icon>
          </v-btn>
        </div>
      </template>
    </v-data-table>
  </div>

</template>
<script>
import ItemListPageBase from '@/components/ItemListPageBase';
import RunnerForm from '@/components/RunnerForm.vue';
import axios from 'axios';

export default {
  components: { RunnerForm },
  mixins: [ItemListPageBase],
  data() {
    return {
      inventories: null,
      projects: null,
    };
  },
  async created() {
    await this.loadData();
  },
  computed: {
    isLoaded() {
      return this.items
        && this.inventories
        && this.projects;
    },
  },
  methods: {
    getHeaders() {
      return [
        {
          text: this.$i18n.t('id'),
          value: 'id',
        },
        {
          text: this.$i18n.t('name'),
          value: 'name',
        },
        {
          text: this.$i18n.t('project'),
          value: 'project',
        },
        {
          text: this.$i18n.t('inventory'),
          value: 'inventory',
        },
        {
          text: this.$i18n.t('webhook'),
          value: 'webhook',
        },
        {
          text: this.$i18n.t('maxParallelTasks'),
          value: 'max_parallel_tasks',
        },
        {
          text: this.$i18n.t('actions'),
          value: 'actions',
          sortable: false,
        },
      ];
    },
    getItemsUrl() {
      return `/api/project/${this.projectId}/runner`;
    },
    getSingleItemUrl() {
      return `/api/project/${this.projectId}/runner/${this.itemId}`;
    },
    getEventName() {
      return 'i-runners';
    },
    async loadData() {
      [this.projects, this.inventories] = (await Promise.all([
        await axios({
          method: 'get',
          url: '/api/projects',
          responseType: 'json',
        }),
        await axios({
          method: 'get',
          url: `/api/project/${this.projectId}/inventory`,
          responseType: 'json',
        }),
      ])).map((x) => x.data);
    },
  },
};
</script>
