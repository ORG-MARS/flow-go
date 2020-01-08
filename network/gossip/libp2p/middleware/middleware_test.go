package middleware

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	mockery "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/network/codec/json"
	"github.com/dapperlabs/flow-go/network/gossip/libp2p/mock"
)

type MiddlewareTestSuit struct {
	suite.Suite
	size int           // used to determine number of middlewares under test
	mws  []*Middleware // used to keep track of middlewares under test
	ov   []*mock.Overlay
	ids  []flow.Identifier
}

// TestMiddlewareTestSuit runs all the test methods in this test suit
func TestMiddlewareTestSuit(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuit))
}

// SetupTest initiates the test setups prior to each test
func (m *MiddlewareTestSuit) SetupTest() {
	m.size = 2 // operates on two middlewares
	m.ids, m.mws = m.createAndStartMiddleWares(m.size)
	require.Len(m.Suite.T(), m.ids, m.size)
	require.Len(m.Suite.T(), m.mws, m.size)
}

func (m *MiddlewareTestSuit) TestSendAndReceive() {
	msg := []byte("hello")
	time.Sleep(4 * time.Second)
	m.mws[0].Send(m.ids[m.size-1], msg)
	time.Sleep(time.Second * 10)
	m.mws[0].Send(m.ids[m.size-1], msg)
	for i := 0; i < m.size; i++ {
		m.ov[i].AssertExpectations(m.T())
	}
}

func (m *MiddlewareTestSuit) createAndStartMiddleWares(count int) ([]flow.Identifier, []*Middleware) {
	var mws []*Middleware
	var ids []flow.Identifier

	// creates the middlewares
	for i := 0; i < count; i++ {
		// generating ids of the nodes
		// as [32]byte{(i+1),0,...,0}
		var target [32]byte
		target[0] = byte(i + 1)
		targetID := flow.Identifier(target)
		ids = append(ids, targetID)

		// generates logger and coder of the nodes
		logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
		codec := json.NewCodec()

		// creates new middleware
		mw, err := New(logger, codec, uint(count-1), "0.0.0.0:0", targetID)
		require.NoError(m.Suite.T(), err)

		mws = append(mws, mw)
	}

	// mocks an overlay (i.e., network) for each middleware
	for i := 0; i < count; i++ {
		overlay := &mock.Overlay{}
		target := i + 1
		if i == count-1 {
			target = 0
		}
		ip, port := mws[target].libP2PNode.GetIPPort()

		// mocks an identity
		flowID := flow.Identity{
			NodeID:  ids[target],
			Address: fmt.Sprintf("%s:%s", ip, port),
			Role:    flow.RoleCollection,
		}

		// mocks Overlay.Identity for middleware.Overlay.Identity()
		overlay.On("Identity").Return(flowID, nil).Once()

		// mocks Overlay.Receive for  middleware.Overlay.Receive(*nodeID, payload)
		overlay.On("Receive", mockery.Anything).Return(nil).Once()
		m.ov = append(m.ov, overlay)
	}

	// starting the middleware
	for i := 0; i < count; i++ {
		mws[i].Start(m.ov[i])
		time.Sleep(1 * time.Second)
	}
	return ids, mws
}
