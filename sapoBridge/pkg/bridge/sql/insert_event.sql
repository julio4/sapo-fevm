INSERT INTO events
	(orderId, attempts, lastAttempt, state, jobSpec, jobId, jobAddr)
    VALUES (:orderId, :attempts, :lastAttempt, :state, :jobSpec, :jobId, :jobAddr);
