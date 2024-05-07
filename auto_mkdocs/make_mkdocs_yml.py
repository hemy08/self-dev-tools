# -*- coding: utf-8 -*-

# 获取指定目录下所有文件列表，并且转成mkdocs中nav的格式

import os
import re
import json
import yaml

# script_path = "F:\\code_hemyzhao\\pages\\hemypages"  # 你的文件路径
script_path = os.getcwd()  # 你的文件路径

script_dir = os.path.basename(script_path)

level = 0

config_file = os.path.join(script_path + "/pages_config.json")

nav_str = ""

JSON_STR_SITE_NAME = "site_name"
JSON_STR_SITE_DESC = "site_description"
JSON_STR_REPO_URL = "repo_url"
JSON_STR_COPYRIGHT = "copyright"
JSON_STR_HOME = "HOME"
JSON_STR_IGNORE_DIRS = "ignore_dirs"
JSON_STR_IGNORE_FILE_TYPES = "ignore_file_types"
JSON_STR_IGNORE_FILES = "ignore_files"
JSON_STR_DOCS_DIR = "docs_dir"
JSON_STR_CHAPTERS = "chapter_infos"
JSON_STR_CHAPTER_RELA_PATH = "relative_path"
JSON_STR_CHAPTER_LINK = "document_link"
JSON_STR_FILE_NAME_CVT = "file_name_convert"
JSON_STR_EXT_LINKS = "external_links"
JSON_STR_THEME = "theme"
JSON_STR_PLUGINS = "plugins"
JSON_STR_EXTRA = "extra"
JSON_STR_EXTRA_CSS = "extra_css"
JSON_STR_EXTRA_JS = "extra_javascript"
JSON_STR_MD_EXTENSIONS = "markdown_extensions"

ignore_dir_list = []
ignore_file_types = []
ignore_files = []
chapters = {}
chapter_links = {}
file_name_cvt_list = {}
external_links = {}
pages_config = {
    JSON_STR_SITE_NAME: '',
    JSON_STR_SITE_DESC: '',
    JSON_STR_REPO_URL: '',
    JSON_STR_COPYRIGHT: '',
    JSON_STR_HOME: '',
    JSON_STR_IGNORE_DIRS: '',
    JSON_STR_IGNORE_FILE_TYPES: '',
    JSON_STR_IGNORE_FILES: '',
    JSON_STR_DOCS_DIR: '',
    JSON_STR_CHAPTERS: '',
    JSON_STR_CHAPTER_RELA_PATH: '',
    JSON_STR_CHAPTER_LINK: '',
    JSON_STR_FILE_NAME_CVT: '',
    JSON_STR_EXT_LINKS: '',
    JSON_STR_THEME: '',
    JSON_STR_PLUGINS: '',
    JSON_STR_EXTRA: '',
    JSON_STR_EXTRA_CSS: '',
    JSON_STR_EXTRA_JS: '',
    JSON_STR_MD_EXTENSIONS: ''
}


# 判断文件路径是否在忽略列表中
def is_ignore(path):
    base = os.path.basename(path)
    return base in ignore_dir_list or base in ignore_files or any(file.endswith(base) for file in ignore_file_types)


# 获取文件列表
def get_file_list(file_path):
    return os.listdir(file_path)


# 文件名转换
def file_name_cvt(file_name):
    return file_name_cvt_list.get(file_name, file_name)


# 删除文件名中的数字前缀
def del_number_prefix(file_name):
    # 如果目录/文件名带数字前缀，去掉 01_、01-、0.1-、0.1_、01 、01.、
    pattern = re.compile("^[0-9.]*[-_ .]")
    match_obj = re.findall(pattern, file_name)
    return file_name[len(match_obj[0]):] if match_obj else file_name


def write_files_info(root_dir, obs_path, name, level, fp):
    docs_dir = pages_config.get(JSON_STR_DOCS_DIR)
    # 去掉.md后缀
    if name.endswith(".md"):
        file_name = name[:len(name) - 3]
        path = obs_path[len(root_dir) + 1 + len(docs_dir) + 1:] if docs_dir else obs_path[len(root_dir) + 1:]
        path = path.replace("\\", "/")
        # 文件名自定义转换
        file_name = file_name_cvt(file_name)

        # 如果目录/文件名带数字前缀，去掉
        file_name = del_number_prefix(file_name)

        # README放在根目录
        if file_name == "README":
            fields = level * " " + "- " + path + "\n"
        else:
            fields = level * " " + "- " + file_name + ": " + path + "\n"
        fp.write(fields)


def print_mkdocs_nav(root_dir, file_dir, level, fp):
    dir_name = os.path.basename(file_dir)
    if level != 0:
        dir_name = del_number_prefix(dir_name)
        cvt_name = file_name_cvt_list.get(dir_name)
        if cvt_name:
            dir_name = cvt_name
        fields = level * " " + "- " + dir_name + ":" + "\n"
        fp.write(fields)

    level = level + 2
    dir_list = get_file_list(file_dir)
    if dir_list is None:
        return

    for name in dir_list:
        if is_ignore(name):
            continue

        # 是一个目录
        obs_path = os.path.abspath(os.path.join(file_dir, name))

        # 根目录的readme文件不加入列表
        if obs_path == os.path.join(root_dir, "README.md"):
            continue

        # 是目录，则递归
        if os.path.isdir(os.path.relpath(obs_path)):
            print_mkdocs_nav(root_dir, obs_path, level, fp)
        else:  # 是一个文件
            write_files_info(root_dir, obs_path, name, level, fp)
    return


def read_pages_conf():
    with open(config_file, 'r', encoding='utf-8') as f:
        res = f.read(-1)
        data = json.loads(res)
        for key in data:
            if key in pages_config:
                pages_config[key] = data[key]
            else:
                pages_config[key] = None
        ignore_dir_list.extend(data[JSON_STR_IGNORE_DIRS].split(";"))
        ignore_files.extend(data[JSON_STR_IGNORE_FILES].split(";"))
        ignore_file_types.extend(data[JSON_STR_IGNORE_FILE_TYPES].split(";"))
        external_links.update(data[JSON_STR_EXT_LINKS])
        chapters.update(data[JSON_STR_CHAPTERS])
        file_name_cvt_list.update(data[JSON_STR_FILE_NAME_CVT])

        f.close()
    return


def write_mkdocs_file(chapter_name, fp):
    root_dir = script_path + chapters.get(chapter_name)
    root_dir = os.path.abspath(root_dir)
    if not os.path.exists(root_dir):
        return

    print_mkdocs_nav(script_path, root_dir, 2, fp)


def write_mkdocs_themes(fp):
    for key in ['theme', 'plugins', 'extra', 'extra_css', 'extra_javascript', 'markdown_extensions']:
        if not pages_config.get(key) is None:
            yaml.dump({key: pages_config.get(key)}, fp, allow_unicode=True)
            fp.write("\n")


def create_chapter_mkdocs():
    yml_file = os.path.join(script_path + "/mkdocs.yml")
    # 打开文件句柄
    with open(yml_file, 'w', encoding='utf-8') as f:
        for key in ['site_name', 'site_description', 'repo_url', 'copyright', 'docs_dir']:
            if not pages_config.get(key) is None and len(pages_config.get(key)):
                yaml.dump({key: pages_config.get(key)}, f, allow_unicode=True)
                f.write("\n")

        write_mkdocs_themes(f)

        f.write("\nnav:\n")

        fields = "  - HOME" + ": " + pages_config.get(JSON_STR_HOME) + "\n"
        f.write(fields)

        for key in chapters:
            write_mkdocs_file(key, f)

        for name in external_links:
            fields = "  - " + name + ": " + external_links.get(name) + "\n"
            f.write(fields)

        f.close()
    return


def mkdocs_mermaid_cvt(file, old_strs, new_strs):
    with open(file, "r", encoding="utf-8") as f:
        file_data = f.read()

    for old, new in zip(old_strs, new_strs):
        if old in file_data:
            file_data = file_data.replace(old, new)

    with open(file, "w", encoding="utf-8") as f:
        f.write(file_data)


old_strs = [
    "format: '!!python/name:pymdownx.superfences.fence_code_format'",
    "emoji_generator: '!!python/name:pymdownx.emoji.to_png'",
    "emoji_generator: '!!python/name:materialx.emoji.to_svg'",
    "emoji_index: '!!python/name:materialx.emoji.twemoji'"
]
new_strs = [
    "format: !!python/name:pymdownx.superfences.fence_code_format",
    "emoji_generator: !!python/name:pymdownx.emoji.to_png",
    "emoji_generator: !!python/name:materialx.emoji.to_svg",
    "emoji_index: !!python/name:materialx.emoji.twemoji"
]

read_pages_conf()
create_chapter_mkdocs()
mkdocs_mermaid_cvt(os.path.join(script_path, "mkdocs.yml"), old_strs, new_strs)
