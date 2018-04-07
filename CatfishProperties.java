package com.td.utils;

import org.json.JSONArray;
import org.json.JSONObject;

import java.io.BufferedReader;

/**
 * Created by Kailash Joshi on 4/6/18.
 */
public class CatfishProperties {
    private BufferedReader reader;
    private String comments = "";
    private Tuple<String, String> prop = new Tuple<>();
    private boolean lastComment = false;
    private boolean lastProp = false;

    public CatfishProperties(BufferedReader reader) {
        this.reader = reader;
    }

    public JSONArray parseProperties() {
        JSONArray array = new JSONArray();
        String st;
        try {
            while ((st = this.reader.readLine()) != null) {

                if (!st.isEmpty()) {
                    if (st.trim().startsWith("#")) {
                        parseComment(st);
                    } else {
                        parseProp(st);
                    }
                }
                if (lastComment & lastProp) {
                    JSONObject obj = new JSONObject();
                    obj.put("Comments", comments.toString());
                    obj.put("Key", prop.getKey());
                    obj.put("Value", prop.getValue());
                    array.put(obj);
                    lastComment = false;
                    lastProp = false;
                    comments = "";
                    prop = new Tuple<>();
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return array;
    }

    private void parseComment(String cmt) {
        if (lastComment) {
            comments = comments + " " + cmt.replace("#", "");
        } else {
            comments = cmt.replace("#", "");
            lastComment = true;
        }
    }

    private void parseProp(String str) {
        if (lastComment == false) {
            String[] tmp = str.split("=", 2);
            prop.add(tmp[0], tmp[1]);
            comments = "Empty comment";
            lastComment = true;
            lastProp = true;
        } else {
            String[] tmp = str.split("=", 2);
            prop.add(tmp[0], tmp[1]);
            lastProp = true;
        }
    }

    class Tuple<K, V> {
        private K k;
        private V v;

        public void add(final K k, final V v) {
            this.k = k;
            this.v = v;
        }

        public K getKey() {
            return k;
        }

        public V getValue() {
            return v;
        }

        @Override
        public String toString() {
            return "Key: " + k + " Value:" + v;
        }
    }
}
