import { metricsActionConstants } from '../constants';

const initialState = {
    metrics: [],
    processing: 0,
};

export function metrics(state = initialState, action) {
    switch (action.type) {
        case metricsActionConstants.GET_METRICS_REQUEST: {
            const processing = state.processing;
            return Object.assign({}, state, {
                processing: processing + 1,
            });
        }

        case metricsActionConstants.GET_METRICS_SUCCESS: {
            const processing = state.processing;
            return Object.assign({}, state, {
                metrics: action.payload,
                processing: processing - 1,
            });
        }

        case metricsActionConstants.GET_METRICS_FAILURE: {
            const processing = state.processing;
            return Object.assign({}, state, {
                metrics: [],
                processing: processing - 1,
            })
        }

        default:
            return state;
    }
}