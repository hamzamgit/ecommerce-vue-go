<template>
        <div class="header-section">
            <img alt="search" src="../assets/search-icon.png" width="20" height="20"/>
            <p class="text-search">Search</p>
            <input type="text" class="search-input" placeholder="Search" v-model="search"/>
        </div>
        <div class="date-section">
            <p class="label">Created date</p>
            <div class="date-input-wrapper">
                <Datepicker
                        range
                        v-model="selectedDate"
                        format="yyyy-dd-MM"
                />
            </div>
        </div>
        <div class="date-section">
            <p class="total-amount">Total amount: <span>$1231.123</span></p>
        </div>
        <div class="table-wrapper">
            <EasyDataTable
                v-model:server-options="serverOptions"
                :headers="headers"
                :items="customerOrders"
                :server-items-length="total"
                hide-footer
            >
                <template #item-orderName="{ orderName, orderId }">
                    <p>{{ orderName }}</p>
                    <p>{{ orderId }}</p>
                </template>
            </EasyDataTable>
        </div>
      <div class="pagination-section">
          <span class="label">Total: {{ total }}</span>
          <select @change="updateRowsPerPageSelect">

              <option
                      v-for="item in rowsPerPageOptions"
                      :value="item"
                      :key="item"
                      :selected="item === rowsPerPageActiveOption"
              >{{ item }}/page</option>
          </select>
          <div class="paginator">
              <div class="icon-wrapper" @click="prevPage()">
                  <img alt="search" class="icon" src="../assets/arrow.png" width="15" height="15" />
              </div>
              <div class="indicator">{{ options.page }}</div>
              <div class="icon-wrapper" @click="nextPage()">
                  <img alt="search" class="icon left" src="../assets/arrow.png" width="15" height="15" />
              </div>
          </div>
      </div>
</template>

<script>
import {mapState, mapActions} from 'vuex'
import Datepicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';
import EasyDataTable from 'vue3-easy-data-table';
import 'vue3-easy-data-table/dist/style.css';


import _ from "lodash"
import {floor} from "lodash/math";

export default {
    name: "CustomerOrders",
    components: {
        EasyDataTable,
        Datepicker
    },
    data() {
        return {
            search: '',
            options: {
                pageSize: 5,
                page: 1,
                sortBy: null,
                sortType: null
            },
            end_date: null,
            start_date: null,
            selectedDate: [],
            rowsPerPageOptions: [5, 50, 100],
            rowsPerPageActiveOption: 5,
            headers: [
                {text: "Order Name", value: "orderName"},
                {text: "Customer Company", value: "customerCompany"},
                {text: "Customer Name", value: "customerName"},
                {text: "Order Date", value: "orderDate", sortable: true},
                {text: "Delivered Amount", value: "deliveredAmount"},
                {text: "Total Amount", value: "totalAmount"},
            ],

        }
    },
    watch: {
        selectedDate: {
            handler(newVal, oldVal) {
                if (newVal) {
                    this.start_date = newVal[0].toISOString()
                    this.end_date = newVal[1].toISOString()
                } else {
                    this.start_date = null
                    this.end_date = null
                }

                this.changeOptions(newVal, oldVal);
            },
            deep: true
        },
        search: {
            handler(newValue, oldValue) {
                this.changeSearch();
            },
        },
        options: {
            handler(newValue, oldValue) {
                this.changeOptions(newValue, oldValue);
            },
            deep: true
        }
    },
    computed: {

        ...mapState({
            customerOrders: state => state.orders.customerOrders,
            total: state => state.orders.total,
        }),
    },
    mounted() {
        this.loadCustomerOrders({
            search: this.search,
            options: this.options,
            start_date: this.start_date,
            end_date: this.end_date
        })
    },
    methods: {
        ...mapActions('orders', [
            'loadCustomerOrders',
        ]),

        changeSearch: _.debounce(function () {
            if (this.options.page != 1) {
                this.options.page = 1;
                // The "options" object will be changed and the "changeOptions" method will be called automatically by "options" watcher
            } else {
                // Call the "changeOptions" method manually
                this.changeOptions();
            }
        }, 1500),

        changeOptions(newValue, oldValue) {
            this.loadCustomerOrders({search: this.search, options: this.options, start_date: this.start_date, end_date: this.end_date})
        },

        nextPage() {
            const check = Math.ceil(this.total / this.options.pageSize)
            if (this.options.page < check) {
                this.options.page += 1
            }
        },

        prevPage() {
            if (this.options.page === 1) {
                this.options.page = 1
            } else {
                this.options.page -= 1
            }
        },

        updateRowsPerPageSelect(event) {
            this.options.pageSize = parseInt(event.target.value)
        }
    }
}
</script>

<style>
.header-section {
    padding: 10px;
    height: 100%;
    display: flex;
    align-items: center;
}

.header-section .text-search {
    margin-left: 10px;
    margin-top: 15px;
    font-weight: 600;
}

.header-section .search-input {
    width: 100%;
    display: flex;
    border: 1px solid #7c7a7d;
    border-radius: 5px;
    margin-left: 20px;
    height: 30px;
    padding-left: 10px;
}

.date-section {
    padding: 10px;
    align-items: center;
    width: max-content;
}

.date-section .label {
    font-weight: bold;
    margin: 0 10px;
}

.date-section .total-amount span {
    font-weight: 600;
}

.date-section .input-wrapper {
    padding: 2px 5px 5px 5px;
    border: 1px solid #7c7a7d;
    border-radius: 5px;
    align-items: center;
}

.date-section .input-wrapper .date-input {
    border: none;
}

.table-wrapper .vue3-easy-data-table[data-v-19cc4b1b] {
    border: none;
}

.date-input-wrapper {
    width: 270px;
}

.pagination-section {
    width: 100%;
    height: 50px;
    align-items: center;
    justify-content: center;
    display: flex;
}

.pagination-section select {
    height: 32px;
}

.pagination-section .label {
    font-size: 15px;
    font-weight: 400;
    color: black;
    margin: 0 10px;
}

.pagination-section .paginator {
    display: flex;
    align-items: center;
    justify-content: center;
    width: max-content;
}

.paginator .icon-wrapper {
    width: 35px;
    height: 35px;
    margin: 0px 10px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.paginator .indicator {
    width: 35px;
    height: 35px;
    align-items: center;
    justify-content: center;
    display: flex;
    border: 1px solid gray;
    border-radius: 5px;
}

.paginator .icon {
    cursor: pointer;
}

.paginator .left {
    transform: rotate(180deg);
}

.pagination-section .indicator-input {
    text-align: center;
    width: 60px;
    height: 25px;
}
</style>