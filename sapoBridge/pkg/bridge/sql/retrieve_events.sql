SELECT eventId, orderId, attempts, lastAttempt, state, jobSpec, jobId, jobAddr
FROM latest_events
WHERE state = :state;
