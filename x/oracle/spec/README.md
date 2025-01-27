# `oracle`

## Overview
The oracle module provides the Provenance Blockchain with the capability to dynamically expose query endpoints through Interchain Queries (ICQ)

One challenge that the Provenance Blockchain faces is supporting each Provenance Blockchain Zone with a unique set of queries. It is not feasible to create an evolving set of queries for each chain. Furthermore, it is not desirable for other parties to request Provenance to build these endpoints for them and then upgrade. This module resolves these issues by enabling Provenance Blockchain zones to manage their own oracle.

## Contents
1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
4. **[Queries](04_queries.md)**
5. **[Events](05_events.md)**
6. **[Genesis](06_genesis.md)**