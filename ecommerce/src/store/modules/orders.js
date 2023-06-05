import axios from 'axios'
import * as types from '../mutation-types'

/**
 * Initial state
 */
export const state = {
    loading: false,
    total: 0,
    customerOrders: [],
    dataOptions: {
        page_size: 5,
        page: 1,
        sort_by: null,
        sort_order: null,
        start_date: null,
        end_date: null,

    },
    last_page: 0,
}


export const actions = {

    async loadCustomerOrders({commit, state}, {search, options, start_date, end_date}) {
        commit(types.SET_DATAOPTIONS, {search: search, options: options, start_date: start_date, end_date: end_date})
        console.log(start_date, end_date)
        commit(types.SET_LOADING, true)
       

console.log(state.dataOptions)
        return await new Promise((resolve, reject) => {
            axios.get('/order', {params: state.dataOptions})
                .then(response => {
                    console.log("REQUEST WAS SUCCESSFUL", response, response.data, response.orders);
                    const data = response.data.orders.map(item => {
                        return {
                            id: item.ID,
                            orderId: item.OrderId,
                            customerCompany: item.User?.Company?.Name,
                            orderName: item.ProductName,
                            customerName: item?.User?.Name,
                            orderDate: item.CreatedAt,
                            deliveredAmount: item.DeliveredAmount,
                            totalAmount: item.TotalAmount
                        }
                    })
                    commit(types.LOAD_CUSTOMER_ORDERS, data)
                    commit(types.SET_TOTAL, response.data.totalCount)
                })
                .catch((error) => {console.log(error)})
                .finally(() => {
                    commit(types.SET_LOADING, false);
                })
        })
    },
}

export const mutations = {

    [types.LOAD_CUSTOMER_ORDERS](state, customerOrders) {
        state.customerOrders = customerOrders
    },

    [types.SET_DATAOPTIONS](state, {search, options, start_date, end_date}) {
        state.dataOptions = {
            search: search,
            sort_by: options.sortBy,
            sort_order: options.sortDesc ? 'desc' : 'asc',
            page_size: options.pageSize,
            page: options.page,
            start_date: start_date,
            end_date: end_date,
        };
    },
    [types.SET_LOADING](state, loading) {
        state.loading = loading
    },

    [types.SET_TOTAL](state, total) {
        state.total = total
    },
}