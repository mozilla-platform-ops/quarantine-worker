License
=======

This Source Code Form is subject to the terms of the Mozilla Public License, v.
2.0. If a copy of the MPL was not distributed with this file, You can obtain
one at http://mozilla.org/MPL/2.0/.

# quarantine-worker
Command line tool for quarantining taskcluster workers

Usage: quarantine-worker provisionerId workerType workerGroup workerId quarantineTime

quarantineTime format: hhmmss. For exemple, if you want to quarantine a worker for 30 days, you will need to provide 720h as value for quarantineTime
