# Beamlight

A lightweight, high-performance Beam Chain client implementation in Go.

## Overview

Beamlight is an implementation of the Ethereum Beam Chain (new name pending) consensus client in Go. The project aims to provide a performant, maintainable, and secure implementation of the Beam Chain specification.

## What is Beam Chain?

Beam Chain is Ethereum's ambitious multi-year initiative to revamp the consensus layer and prepare Ethereum's foundations for the next century. Key components include:

- **Attester-Proposer Separation (APS)**: Separation of validator duties to enhance security and efficiency
- **3-Slot Finality (3SF)**: Faster finality with a simplified protocol
- **Post-Quantum Security**: Transition to post-quantum cryptography using hash-based signatures
- **Chain Snarkification**: Enabling proof generation for consensus validation
- **P2P Networking Improvements**: Enhanced networking with generalized gossipsub, grid topology, and more

## Project Structure

```
beamlight/
├── cmd/                      # Command-line applications
│   └── beamlight/            # Main application entry point
├── internal/                 # Private application code
│   ├── beacon/               # Core beacon chain functionality
│   ├── config/               # Configuration handling
│   │   ├── flags/            # CLI flags handling
│   │   └── params/           # Network parameters
│   ├── consensus/            # Consensus implementation
│   ├── crypto/               # Cryptographic primitives
│   ├── database/             # Database abstractions
│   ├── node/                 # Node service orchestration
│   ├── p2p/                  # P2P networking
│   ├── rpc/                  # RPC implementations
│   └── zk/                   # ZK-related functionality
│       └── stf/              # State transition function (ZK-friendly)
├── pkg/                      # Public libraries for external use
├── api/                      # API definitions
├── build/                    # Build-related files
├── docs/                     # Documentation
├── scripts/                  # Helper scripts
├── test/                     # Tests
└── tools/                    # Development tools
```

## Goals

- Create a production-ready Beam Chain client in Go
- Optimize for performance and resource usage
- Implement all major Beam Chain features
- Contribute to the Ethereum ecosystem by increasing client diversity
- Focus on developer ergonomics and maintainability

## Roadmap

### Phase 1: Bootstrap (Q2-Q3 2025)
- [ ] Set up project structure and tooling
- [ ] Implement basic network connectivity
- [ ] Implement basic state management
- [ ] Participate in testnet

### Phase 2: Core Implementation (Q3-Q4 2025)
- [ ] Implement 3SF consensus
- [ ] Implement APS
- [ ] Implement P2P networking improvements

### Phase 3: Optimization (Q1 2026)
- [ ] Optimize performance
- [ ] Implement post-quantum signatures
- [ ] Prepare for chain snarkification

### Phase 4: Production Release (Q2-Q3 2026)
- [ ] Security audits
- [ ] Mainnet readiness
- [ ] Production deployment

## Technical Considerations

### Post-Quantum Cryptography

Beamlight will implement hash-based signatures as part of Ethereum's shift towards post-quantum security. This includes:

- Winternitz XMSS signatures as BLS replacement
- Efficient implementation of hash-based signature aggregation

### ZK-VM Integration

To support chain snarkification, Beamlight will need to:

- Separate state transition function (STF) for easy snarkification
- Minimize STF dependencies
- Implement stateless reads and writes for efficient proof generation

### P2P Networking

Beamlight will implement improved P2P networking with:

- Generalized gossipsub
- Gossipsub v2
- Grid topology for improved network resilience
- Efficient set reconciliation

## Contribution

Beamlight is an open-source project and welcomes contributions in various forms:

- Code contributions
- Documentation
- Testing
- Feature suggestions
- Bug reports

## Resources

### Official Beam Chain Resources
- [Beam Roadmap](https://beamroadmap.org/)
- [Beam Call Recordings](https://beamroadmap.org/calls)
- [Ethereum Research on Beam Chain](https://ethresear.ch/)

### Development Resources
- [Ethereum Consensus Specs](https://github.com/ethereum/consensus-specs)
- [Go Ethereum](https://github.com/ethereum/go-ethereum)

## License

Beamlight is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html). 