import React, { useState, useEffect } from "react";
import { Tabs, message, Row, Col } from "antd";
import axios from "axios";

import SearchBar from "./SearchBar";
import PhotoGallery from "./PhotoGallery";
import CreatePostButton from "./CreatePostButton";
import ImageGenerator from "./ImageGenerator"; // Import the ImageGenerator component
import { SEARCH_KEY, BASE_URL, TOKEN_KEY } from "../constants";

const { TabPane } = Tabs;

function Home(props) {
    const [posts, setPost] = useState([]);
    const [activeTab, setActiveTab] = useState("image");
    const [searchOption, setSearchOption] = useState({
        type: SEARCH_KEY.all,
        keyword: ""
    });

    const handleSearch = (option) => {
        const { type, keyword } = option;
        setSearchOption({ type: type, keyword: keyword });
    };

    useEffect(() => {
        fetchPost(searchOption);
    }, [searchOption]);

    const fetchPost = (option) => {
        const { type, keyword } = option;
        let url = "";
    
        if (type === SEARCH_KEY.all) {
            url = `${BASE_URL}/search`;
        } else if (type === SEARCH_KEY.user) {
            url = `${BASE_URL}/search?user=${keyword}`;
        } else {
            url = `${BASE_URL}/search?keywords=${keyword}`;
        }
    
        const opt = {
            method: "GET",
            url: url,
            headers: {
                Authorization: `Bearer ${localStorage.getItem(TOKEN_KEY)}`
            }
        };
    
        axios(opt)
            .then((res) => {
                if (res.status === 200) {
                    console.log("Fetched posts: ", res.data); // Log the fetched posts
                    setPost(res.data);
                }
            })
            .catch((err) => {
                message.error("Fetch posts failed!");
                console.log("fetch posts failed: ", err.message);
            });
    };
    

    const renderPosts = (type) => {
        if (!posts || posts.length === 0) {
            return <div>No data!</div>;
        }
    
        if (type === "image") {
            const imageArr = posts
                .filter((item) => item.type === "image")
                .map((image) => {
                    return {
                        postId: image.id,
                        src: image.url,
                        user: image.user,
                        caption: image.message,
                        thumbnail: image.url,
                        thumbnailWidth: 300,
                        thumbnailHeight: 200
                    };
                });
    
            return <PhotoGallery images={imageArr} />;
        } else if (type === "video") {
            return (
                <div style={{ maxHeight: "900px" }}>
                    <Row gutter={[1, 1]}>
                        {posts
                            .filter((post) => post.type === "video")
                            .map((post) => (
                                <Col span={24} key={post.id}>
                                    <video 
                                        src={post.url} 
                                        controls={true} 
                                        style={{ width: "100%" }} 
                                    />
                                    <p>{post.user}: {post.message}</p>
                                </Col>
                            ))}
                    </Row>
                </div>
            );
        }
    };
    
    

    const showPost = (type) => {
        setActiveTab(type);
        setTimeout(() => {
            setSearchOption({ type: SEARCH_KEY.all, keyword: "" });
        }, 3000);
    };

    return (
        <div className="home">
            <SearchBar handleSearch={handleSearch} />
            <div className="display">
                <Tabs
                    onChange={(key) => setActiveTab(key)}
                    defaultActiveKey="image"
                    activeKey={activeTab}
                    tabBarExtraContent={
                        <div style={{ display: 'flex', gap: '10px' }}>
                            <ImageGenerator /> {/* Keep this functional button */}
                            <CreatePostButton onShowPost={showPost} />
                        </div>
                    }
                >
                    <TabPane tab="Images" key="image">
                        {renderPosts("image")}
                    </TabPane>
                    <TabPane tab="Videos" key="video">
                        {renderPosts("video")}
                    </TabPane>
                </Tabs>
            </div>
        </div>
    );
}

export default Home;