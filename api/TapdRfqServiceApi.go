// //go:build no_grpc
package api

//
//import (
//	"context"
//	"fmt"
//	"github.com/lightninglabs/taproot-assets/taprpc/rfqrpc"
//	"github.com/wallet/api/connect"
//	"github.com/wallet/base"
//	"google.golang.org/grpc"
//	"io"
//	"path/filepath"
//)
//
//// AddAssetBuyOrder
////
////	@Description:AddAssetBuyOrder is used to add a buy order for a specific asset. If a buy order already exists for the asset, it will be updated.
////	@return bool
//func AddAssetBuyOrder() bool {
//	grpcHost := base.QueryConfigByKey("taproothost")
//	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
//	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), base.NetWork)
//	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
//	creds := connect.NewTlsCert(tlsCertPath)
//	macaroon := connect.GetMacaroon(macaroonPath)
//	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
//		grpc.WithPerRPCCredentials(api.NewMacaroonCredential(macaroon)))
//	if err != nil {
//		fmt.Printf("%s did not connect: grpc.Dial: %v\n", api.GetTimeNow(), err)
//	}
//	defer func(conn *grpc.ClientConn) {
//		err := conn.Close()
//		if err != nil {
//			fmt.Printf("%s conn Close Error: %v\n", api.GetTimeNow(), err)
//		}
//	}(conn)
//	client := rfqrpc.NewRfqClient(conn)
//	request := &rfqrpc.AddAssetBuyOrderRequest{}
//	response, err := client.AddAssetBuyOrder(context.Background(), request)
//	if err != nil {
//		fmt.Printf("%s rfqrpc AddAssetBuyOrder Error: %v\n", api.GetTimeNow(), err)
//		return false
//	}
//	fmt.Printf("%s %v\n", api.GetTimeNow(), response)
//	return true
//}
//
//// AddAssetSellOffer
////
////	@Description:AddAssetSellOffer is used to add a sell offer for a specific asset. If a sell offer already exists for the asset, it will be updated.
////	@return bool
//func AddAssetSellOffer() bool {
//	grpcHost := base.QueryConfigByKey("taproothost")
//	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
//	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), base.NetWork)
//	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
//	creds := connect.NewTlsCert(tlsCertPath)
//	macaroon := connect.GetMacaroon(macaroonPath)
//	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
//		grpc.WithPerRPCCredentials(api.NewMacaroonCredential(macaroon)))
//	if err != nil {
//		fmt.Printf("%s did not connect: grpc.Dial: %v\n", api.GetTimeNow(), err)
//	}
//	defer func(conn *grpc.ClientConn) {
//		err := conn.Close()
//		if err != nil {
//			fmt.Printf("%s conn Close Error: %v\n", api.GetTimeNow(), err)
//		}
//	}(conn)
//	client := rfqrpc.NewRfqClient(conn)
//	request := &rfqrpc.AddAssetSellOfferRequest{}
//	response, err := client.AddAssetSellOffer(context.Background(), request)
//	if err != nil {
//		fmt.Printf("%s rfqrpc  Error: %v\n", api.GetTimeNow(), err)
//		return false
//	}
//	fmt.Printf("%s %v\n", api.GetTimeNow(), response)
//	return true
//}
//
//// QueryRfqAcceptedQuotes
////
////	@Description:
////	@return string
//func QueryRfqAcceptedQuotes() string {
//	grpcHost := base.QueryConfigByKey("taproothost")
//	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
//	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), base.NetWork)
//	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
//	creds := connect.NewTlsCert(tlsCertPath)
//	macaroon := connect.GetMacaroon(macaroonPath)
//	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
//		grpc.WithPerRPCCredentials(api.NewMacaroonCredential(macaroon)))
//	if err != nil {
//		fmt.Printf("%s did not connect: grpc.Dial: %v\n", api.GetTimeNow(), err)
//	}
//	defer func(conn *grpc.ClientConn) {
//		err := conn.Close()
//		if err != nil {
//			fmt.Printf("%s conn Close Error: %v\n", api.GetTimeNow(), err)
//		}
//	}(conn)
//	client := rfqrpc.NewRfqClient(conn)
//	request := &rfqrpc.QueryPeerAcceptedQuotesRequest{}
//	response, err := client.QueryPeerAcceptedQuotes(context.Background(), request)
//	if err != nil {
//		fmt.Printf("%s rfqrpc QueryRfqAcceptedQuotes Error: %v\n", api.GetTimeNow(), err)
//		return ""
//	}
//	return response.String()
//}
//
//// SubscribeRfqEventNtfns
////
////	@Description:SubscribeRfqEventNtfns is used to subscribe to RFQ events.
////	@return bool
//func SubscribeRfqEventNtfns() bool {
//	grpcHost := base.QueryConfigByKey("taproothost")
//	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
//	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), base.NetWork)
//	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
//	creds := connect.NewTlsCert(tlsCertPath)
//	macaroon := connect.GetMacaroon(macaroonPath)
//	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
//		grpc.WithPerRPCCredentials(api.NewMacaroonCredential(macaroon)))
//	if err != nil {
//		fmt.Printf("%s did not connect: grpc.Dial: %v\n", api.GetTimeNow(), err)
//	}
//	defer func(conn *grpc.ClientConn) {
//		err := conn.Close()
//		if err != nil {
//			fmt.Printf("%s conn Close Error: %v\n", api.GetTimeNow(), err)
//		}
//	}(conn)
//	client := rfqrpc.NewRfqClient(conn)
//	request := &rfqrpc.SubscribeRfqEventNtfnsRequest{}
//	stream, err := client.SubscribeRfqEventNtfns(context.Background(), request)
//	if err != nil {
//		fmt.Printf("%s rfqrpc  Error: %v\n", api.GetTimeNow(), err)
//		return false
//	}
//	for {
//		response, err := stream.Recv()
//		if err != nil {
//			if err == io.EOF {
//				fmt.Printf("%s err == io.EOF, err: %v\n", api.GetTimeNow(), err)
//				return false
//			}
//			fmt.Printf("%s stream Recv err: %v\n", api.GetTimeNow(), err)
//			return false
//		}
//		fmt.Printf("%s %v\n", api.GetTimeNow(), response)
//		return true
//	}
//
//}
